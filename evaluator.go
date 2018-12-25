package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/olekukonko/tablewriter"
	"github.com/tzmfreedom/go-soapforce"
)

type Evaluator interface {
	Select(string) error
	Insert(*Statement) error
	Update(*Statement) error
	Delete(*Statement) error
}

type DefaultEvaluator struct {
	Client *soapforce.Client
}

func (e *DefaultEvaluator) Select(q string) error {
	q, err := expandWildcard(q, e.Client)
	if err != nil {
		return err
	}
	r, err := e.Client.Query(q)
	if err != nil {
		return err
	}
	fieldNames, err := getFields(q)
	if err != nil {
		return err
	}
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader(fieldNames)

	for _, record := range r.Records {
		fields := map[string]interface{}{}
		for f, v := range record.Fields {
			fields[strings.ToUpper(f)] = v
		}
		columns := []string{}
		for _, fieldName := range fieldNames {
			if fieldName == "ID" {
				columns = append(columns, record.Id)
			} else {
				value := ""
				if v := fields[fieldName]; v != nil {
					value = v.(string)
				}
				columns = append(columns, value)
			}
		}
		table.Append(columns)
	}
	table.Render()
	return nil
}

func (e *DefaultEvaluator) Insert(stmt *Statement) error {
	obj := &soapforce.SObject{}
	obj.Type = stmt.Sobject
	obj.Fields = map[string]interface{}{}
	for field, value := range stmt.Values {
		obj.Fields[field] = value
	}
	results, err := e.Client.Create([]*soapforce.SObject{obj})
	if err != nil {
		return err
	}
	res := results[0]
	if res.Success {
		fmt.Printf("%s Created: Id = %s\n", stmt.Sobject, res.Id)
	} else {
		for _, error := range res.Errors {
			fmt.Fprintln(os.Stderr, error.Message)
		}
	}
	return nil
}

func (e *DefaultEvaluator) Update(stmt *Statement) error {
	q := fmt.Sprintf("SELECT id FROM %s %s", stmt.Sobject, stmt.Where)
	r, err := e.Client.Query(q)
	if err != nil {
		return err
	}
	updateObjects := make([]*soapforce.SObject, len(r.Records))
	for i, record := range r.Records {
		obj := &soapforce.SObject{}
		obj.Id = record.Id
		obj.Type = stmt.Sobject
		obj.Fields = map[string]interface{}{}
		for field, value := range stmt.Values {
			obj.Fields[field] = value
		}
		updateObjects[i] = obj
	}
	results, err := e.Client.Update(updateObjects)
	if err != nil {
		return err
	}
	for _, res := range results {
		if res.Success {
			fmt.Printf("%s Updated: Id = %s\n", stmt.Sobject, res.Id)
		} else {
			for _, error := range res.Errors {
				fmt.Fprintln(os.Stderr, error.Message)
			}
		}
	}
	return nil
}

func (e *DefaultEvaluator) Delete(stmt *Statement) error {
	q := fmt.Sprintf("SELECT id FROM %s %s", stmt.Sobject, stmt.Where)
	r, err := e.Client.Query(q)
	if err != nil {
		return err
	}
	ids := make([]string, len(r.Records))
	for i, record := range r.Records {
		ids[i] = record.Id
	}
	results, err := e.Client.Delete(ids)
	if err != nil {
		return err
	}
	for _, res := range results {
		if res.Success {
			fmt.Printf("%s Deleted: Id = %s\n", stmt.Sobject, res.Id)
		} else {
			for _, error := range res.Errors {
				fmt.Fprintln(os.Stderr, error.Message)
			}
		}
	}
	return nil
}

type DryRunner struct {
	Client *soapforce.Client
}

func (e *DryRunner) Select(q string) error {
	q, err := expandWildcard(q, e.Client)
	if err != nil {
		return err
	}
	r, err := e.Client.Query(q)
	if err != nil {
		return err
	}
	fieldNames, err := getFields(q)
	if err != nil {
		return err
	}
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader(fieldNames)

	for _, record := range r.Records {
		fields := map[string]interface{}{}
		for f, v := range record.Fields {
			fields[strings.ToUpper(f)] = v
		}
		columns := []string{}
		for _, fieldName := range fieldNames {
			if fieldName == "ID" {
				columns = append(columns, record.Id)
			} else {
				columns = append(columns, fields[fieldName].(string))
			}
		}
		table.Append(columns)
	}
	table.Render()
	return nil
}

func (e *DryRunner) Insert(stmt *Statement) error {
	fields := make([]string, len(stmt.Values))
	i := 0
	for field, value := range stmt.Values {
		fields[i] = fmt.Sprintf("%s = %s", field, value)
		i++
	}
	fmt.Printf("%s Created: %s\n", stmt.Sobject, strings.Join(fields, ", "))
	return nil
}

func (e *DryRunner) Update(stmt *Statement) error {
	q := fmt.Sprintf("SELECT id FROM %s %s", stmt.Sobject, stmt.Where)
	r, err := e.Client.Query(q)
	if err != nil {
		return err
	}

	for _, record := range r.Records {
		fields := make([]string, len(stmt.Values))
		i := 0
		for field, value := range stmt.Values {
			fields[i] = fmt.Sprintf("%s = %s", field, value)
			i++
		}
		fmt.Printf("%s Updated: Id = %s, %s\n", stmt.Sobject, record.Id, strings.Join(fields, ", "))
	}
	return nil
}

func (e *DryRunner) Delete(stmt *Statement) error {
	q := fmt.Sprintf("SELECT id FROM %s %s", stmt.Sobject, stmt.Where)
	r, err := e.Client.Query(q)
	if err != nil {
		return err
	}
	ids := make([]string, len(r.Records))
	for i, record := range r.Records {
		ids[i] = record.Id
		fmt.Printf("%s Deleted: Id = %s\n", stmt.Sobject, record.Id)
	}
	return nil
}

func getFields(q string) ([]string, error) {
	r, err := regexp.Compile(`SELECT\s+(.*)\s+FROM\s.*`)
	if err != nil {
		return nil, err
	}
	matches := r.FindStringSubmatch(q)
	if len(matches) == 1 {
		return []string{}, nil
	}
	fields := strings.Split(matches[1], ",")
	trimmedFields := make([]string, len(fields))
	for i, field := range fields {
		trimmedFields[i] = strings.ToUpper(strings.TrimSpace(field))
	}
	return trimmedFields, nil
}

func expandWildcard(original string, c *soapforce.Client) (string, error) {
	r := regexp.MustCompile(`(?i)SELECT\s+(\*)\s+FROM\s+([a-zA-Z\d_]+)`)
	matches := r.FindStringSubmatch(strings.TrimSpace(original))
	if len(matches) == 0 {
		return original, nil
	}
	result, err := c.DescribeSObject(matches[2])
	if err != nil {
		return "", err
	}
	fields := make([]string, len(result.Fields))
	for i, f := range result.Fields {
		fields[i] = f.Name
	}
	selectClause := strings.Join(fields, ",")
	return r.ReplaceAllString(original, fmt.Sprintf("SELECT %s FROM $2", selectClause)), nil
}
