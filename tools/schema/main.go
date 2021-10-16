// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/iotaledger/wasp/tools/schema/generator"
	"gopkg.in/yaml.v2"
)

var (
	disabledFlag   = false
	flagCore       = flag.Bool("core", false, "generate core contract interface")
	flagForce      = flag.Bool("force", false, "force code generation")
	flagGo         = flag.Bool("go", false, "generate Go code")
	flagInit       = flag.String("init", "", "generate new schema file for smart contract named <string>")
	flagJava       = &disabledFlag // flag.Bool("java", false, "generate Java code <outdated>")
	flagRust       = flag.Bool("rust", false, "generate Rust code <default>")
	flagSchemaType = flag.String("schemaType", "json", "Extension of the schema that will be generated/used. Values(json,yaml)")
)

const (
	schema = "schema."
)

func init() {
	flag.Parse()
}

func main() {
	err := generator.FindModulePath()
	if err != nil {
		fmt.Println(err)
		return
	}

	file, err := os.Open(schema + *flagSchemaType)
	if err == nil {
		defer file.Close()
		if *flagInit != "" {
			fmt.Println(schema + *flagSchemaType + " already exists")
			return
		}
		err = generateSchema(file)
		if err != nil {
			fmt.Println(err)
		}
		return
	}

	if *flagInit != "" {
		err = generateSchemaNew(*flagSchemaType)
		if err != nil {
			fmt.Println(err)
		}
		return
	}

	flag.Usage()
}

func generateSchema(file *os.File) error {
	info, err := file.Stat()
	if err != nil {
		return err
	}
	schemaTime := info.ModTime()

	schema, err := loadSchema(file)
	if err != nil {
		return err
	}

	schema.CoreContracts = *flagCore
	if *flagGo {
		info, err = os.Stat("consts.go")
		if err == nil && info.ModTime().After(schemaTime) && !*flagForce {
			fmt.Println("skipping Go code generation")
		} else {
			fmt.Println("generating Go code")
			err = schema.GenerateGo()
			if err != nil {
				return err
			}
			if !schema.CoreContracts {
				err = schema.GenerateGoTests()
				if err != nil {
					return err
				}
			}
		}
	}

	if *flagJava {
		fmt.Println("generating Java code")
		err = schema.GenerateJava()
		if err != nil {
			return err
		}
	}

	if *flagRust {
		info, err = os.Stat("src/consts.rs")
		if err == nil && info.ModTime().After(schemaTime) && !*flagForce {
			fmt.Println("skipping Rust code generation")
		} else {
			fmt.Println("generating Rust code")
			err = schema.GenerateRust()
			if err != nil {
				return err
			}
			if !schema.CoreContracts {
				err = schema.GenerateGoTests()
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func generateSchemaNew(schemaExtension string) error {
	name := *flagInit
	fmt.Println("initializing " + name)

	subfolder := strings.ToLower(name)
	err := os.Mkdir(subfolder, 0o755)
	if err != nil {
		return err
	}
	err = os.Chdir(subfolder)
	if err != nil {
		return err
	}

	file, err := os.Create(schema + schemaExtension)
	if err != nil {
		return err
	}
	defer file.Close()

	templateSchema := &generator.TemplateSchema{}
	templateSchema.Name = name
	templateSchema.Description = name + " description"
	templateSchema.Structs = make(generator.StringMapMap)
	templateSchema.Typedefs = make(generator.StringMap)
	templateSchema.State = make(generator.StringMap)
	templateSchema.State["owner"] = "AgentID // current owner of this smart contract"
	templateSchema.Funcs = make(generator.FuncDescMap)
	templateSchema.Views = make(generator.FuncDescMap)

	funcInit := &generator.FuncDesc{}
	funcInit.Params = make(generator.StringMap)
	funcInit.Params["owner"] = "?AgentID // optional owner of this smart contract"
	templateSchema.Funcs["init"] = funcInit

	funcSetOwner := &generator.FuncDesc{}
	funcSetOwner.Access = "owner // current owner of this smart contract"
	funcSetOwner.Params = make(generator.StringMap)
	funcSetOwner.Params["owner"] = "AgentID // new owner of this smart contract"
	templateSchema.Funcs["setOwner"] = funcSetOwner

	viewGetOwner := &generator.FuncDesc{}
	viewGetOwner.Results = make(generator.StringMap)
	viewGetOwner.Results["owner"] = "AgentID // current owner of this smart contract"
	templateSchema.Views["getOwner"] = viewGetOwner

	if schemaExtension == "json" {
		err := WriteJsonSchema(templateSchema, file)
		return err
	} else if schemaExtension == "yaml" {
		err := WriteYamlSchema(templateSchema, file)
		return err
	}
	return errors.New("Not implemented Schema Type " + schemaExtension)
}

func WriteJsonSchema(templateSchema *generator.TemplateSchema, file *os.File) error {
	b, err := json.Marshal(templateSchema)
	if err != nil {
		return err
	}

	var out bytes.Buffer
	err = json.Indent(&out, b, "", "\t")
	if err != nil {
		return err
	}

	_, err = out.WriteTo(file)
	return err
}

func WriteYamlSchema(templateSchema *generator.TemplateSchema, file *os.File) error {
	b, err := yaml.Marshal(templateSchema)
	if err != nil {
		return err
	}
	_, err = file.Write(b)
	return err
}

func loadSchema(file *os.File) (*generator.Schema, error) {
	fmt.Println("loading " + file.Name())
	fileExtension := filepath.Ext(file.Name())
	templateSchema := &generator.TemplateSchema{}

	if fileExtension == ".json" {
		err := json.NewDecoder(file).Decode(templateSchema)
		if err != nil {
			return nil, err
		}
	} else if fileExtension == ".yaml" {
		fileByteArray, _ := ioutil.ReadAll(file)
		err := yaml.Unmarshal(fileByteArray, templateSchema)
		if err != nil {
			return nil, err
		}
	} else {
		return nil, errors.New("Not implemented Schema Type " + fileExtension)
	}

	schema := generator.NewSchema()
	err := schema.Compile(templateSchema)
	if err != nil {
		return nil, err
	}
	return schema, nil
}
