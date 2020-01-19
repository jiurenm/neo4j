package dao

import (
	"fmt"
	bolt "github.com/johnnadratowski/golang-neo4j-bolt-driver"
	"neo4j/pkg/conf"
	"strings"
)

type DB struct {
	neo4j bolt.Conn
}

func New(configs *conf.Configs) (*DB, error) {
	driver := bolt.NewDriver()
	uri := fmt.Sprintf("bolt://%s:%s@%s:%s", configs.Neo4j.Username, configs.Neo4j.Password, configs.Neo4j.Host, configs.Neo4j.Port)
	conn, err := driver.OpenNeo(uri)
	if err != nil {
		return nil, err
	}
	return &DB{neo4j: conn}, nil
}

func (db *DB) FindByName(name string) (string, error) {
	data, _, _, err := db.neo4j.QueryNeoAll(
		"MATCH (n:disease) WHERE n.name =~ {name} RETURN n.info", map[string]interface{}{"name": name})
	if err != nil {
		return "", err
	}
	if len(data) > 0 {
		var list = make([]string, 0)
		for _, value := range data {
			list = append(list, value[0].(string))
		}
		reverse(list)
		return strings.Join(list, ""), nil
	}
	return "", nil
}

func (db *DB) FindName() ([]string, error) {
	data, _, _, err := db.neo4j.QueryNeoAll(
		"MATCH (n:disease) WHERE n.name =~ {name} RETURN n.name", map[string]interface{}{"name": ".*"})
	if err != nil {
		return nil, err
	}
	var list = make([]string, 0)
	for _, value := range data {
		list = append(list, value[0].(string))
	}
	reverse(list)
	return list, nil
}

func (db *DB) FindCauseByName(name string) (string, error) {
	data, _, _, err := db.neo4j.QueryNeoAll(
		"MATCH (n:disease) - [r:CAUSED] -> (m:cause) WHERE n.name =~ {name} RETURN m.info", map[string]interface{}{"name": name})
	if err != nil {
		return "", err
	}
	if len(data) > 0 {
		var list = make([]string, 0)
		for _, value := range data {
			list = append(list, value[0].(string))
		}
		reverse(list)
		return strings.Join(list, ""), nil
	}
	return "", nil
}

func (db *DB) FindDiagnosisByName(name string) (string, error) {
	data, _, _, err := db.neo4j.QueryNeoAll(
		"MATCH (n:disease) - [r:DIAGNOSED] -> (m:diagnosis) WHERE n.name =~ {name} RETURN m.info", map[string]interface{}{"name": name})
	if err != nil {
		return "", err
	}
	if len(data) > 0 {
		var list = make([]string, 0)
		for _, value := range data {
			list = append(list, value[0].(string))
		}
		reverse(list)
		return strings.Join(list, ""), nil
	}
	return "", nil
}

func (db *DB) FindHarmByName(name string) (string, error) {
	data, _, _, err := db.neo4j.QueryNeoAll(
		"MATCH (n:disease) - [r:HAS_HARM] -> (m:harm) WHERE n.name =~ {name} RETURN m.info", map[string]interface{}{"name": name})
	if err != nil {
		return "", err
	}
	if len(data) > 0 {
		var list = make([]string, 0)
		for _, value := range data {
			list = append(list, value[0].(string))
		}
		reverse(list)
		return strings.Join(list, ""), nil
	}
	return "", nil
}

func (db *DB) FindNoteByName(name string) (string, error) {
	data, _, _, err := db.neo4j.QueryNeoAll(
		"MATCH (n:disease) - [r:HAS_NOTE] -> (m:note) WHERE n.name =~ {name} RETURN m.info", map[string]interface{}{"name": name})
	if err != nil {
		return "", err
	}
	if len(data) > 0 {
		var list = make([]string, 0)
		for _, value := range data {
			list = append(list, value[0].(string))
		}
		reverse(list)
		return strings.Join(list, ""), nil
	}
	return "", nil
}

func (db *DB) FindPreventionByName(name string) (string, error) {
	data, _, _, err := db.neo4j.QueryNeoAll(
		"MATCH (n:disease) - [r:PREVENTED] -> (m:prevention) WHERE n.name =~ {name} RETURN m.info", map[string]interface{}{"name": name})
	if err != nil {
		return "", err
	}
	if len(data) > 0 {
		var list = make([]string, 0)
		for _, value := range data {
			list = append(list, value[0].(string))
		}
		reverse(list)
		return strings.Join(list, ""), nil
	}
	return "", nil
}

func (db *DB) FindSymptomByName(name string) (string, error) {
	data, _, _, err := db.neo4j.QueryNeoAll(
		"MATCH (n:disease) - [r:HAS_SYMPTOM] -> (m:symptom) WHERE n.name =~ {name} RETURN m.info", map[string]interface{}{"name": name})
	if err != nil {
		return "", err
	}
	if len(data) > 0 {
		var list = make([]string, 0)
		for _, value := range data {
			list = append(list, value[0].(string))
		}
		reverse(list)
		return strings.Join(list, ""), nil
	}
	return "", nil
}

func (db *DB) FindTreatmentByName(name string) (string, error) {
	data, _, _, err := db.neo4j.QueryNeoAll(
		"MATCH (n:disease) - [r:TREATED] -> (m:treatment) WHERE n.name =~ {name} RETURN m.info", map[string]interface{}{"name": name})
	if err != nil {
		return "", err
	}
	if len(data) > 0 {
		var list = make([]string, 0)
		for _, value := range data {
			list = append(list, value[0].(string))
		}
		reverse(list)
		return strings.Join(list, ""), nil
	}
	return "", nil
}

func reverse(s []string) []string {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}
