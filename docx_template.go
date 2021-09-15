package main

import (
	"archive/zip"
	"io"
	"log"

	"strings"

	"github.com/beevik/etree"
)

const DOCUMENT_CONTENT = "word/document.xml"

type DocxTemplate struct {
	source  *zip.ReadCloser
	content *etree.Document //"word/document.xml"
}

func OpenTemplate(path string) (DocxTemplate, error) {
	var docx DocxTemplate
	var err error
	docx.source, err = zip.OpenReader(path)
	if err != nil {
		return docx, err
	}
	for _, f := range docx.source.File {
		if f.Name != DOCUMENT_CONTENT {
			continue
		}
		rc, err := f.Open()
		if err != nil {
			return docx, err
		}
		docx.content = etree.NewDocument()
		docx.content.ReadFrom(rc)
		rc.Close()
	}
	return docx, nil
}

func (docx *DocxTemplate) Close() error {
	return docx.source.Close()
}

func (docx *DocxTemplate) Write(output io.Writer) error {
	writer := zip.NewWriter(output)
	for _, fsource := range docx.source.File {
		fout, err := writer.Create(fsource.Name)
		if err != nil {
			return err
		}
		if fsource.Name == DOCUMENT_CONTENT {
			docx.content.WriteTo(fout)
		} else {
			fin, err := fsource.Open()
			if err != nil {
				return err
			}
			defer fin.Close()
			_, err = io.Copy(fout, fin)
			if err != nil {
				return err
			}
		}
	}
	err := writer.Close()
	if err != nil {
		return err
	}
	return nil
}

type TreeWalkFunction func(e *etree.Element)

func walkTree(root *etree.Element, callback TreeWalkFunction) {
	for _, r := range root.ChildElements() {
		callback(r)
		walkTree(r, callback)
	}
}

//walk tree, replace text in <t> elements
//replacementPairs = []string{"#template_var1#", "value1", "#template_var2#", "value2"...}
func (docx *DocxTemplate) Replace(replacementPairs []string) {
	replacer := strings.NewReplacer(replacementPairs...)
	walkTree(docx.content.Root(), func(e *etree.Element) {
		if e.Tag == "t" {
			e.SetText(replacer.Replace(e.Text()))
		}
	})
}

func (docx *DocxTemplate) ReplaceToList(parent_tag string, marker string, replacementList [][]string) {
	var multilineT []*etree.Element = make([]*etree.Element, 0)
	var multilineP *etree.Element = nil
	var multilineParent *etree.Element = nil
	var multilinePos int = -1
	//1. находим все <t>, содержащий #marker#
	walkTree(docx.content.Root(), func(e *etree.Element) {
		if e.Tag == "t" {
			if strings.Index(e.Text(), marker) >= 0 {
				multilineT = append(multilineT, e)
			}
		}
	})
	if len(multilineT) == 0 {
		return
	}
	for _, mT := range multilineT {
		//2. находим ближайшего предка <t> с тегом parent_tag
		multilineP = mT.Parent()
		for {
			if multilineP == nil {
				log.Printf("ReplaceToList: (%v) no parent (multilineP)", mT)
				goto CONT
			}
			if multilineP.Tag == parent_tag {
				break
			}
			multilineP = multilineP.Parent()
		}
		multilineParent = multilineP.Parent()
		if multilineParent == nil {
			log.Printf("ReplaceToList: (%v) no parent (multilineParent)", mT)
			goto CONT
		}
		multilinePos = multilineP.Index()
		//создаем копии parent_tag, заменяем в них поля, вставляем перед parent_tag
		for _, replPairs := range replacementList {
			copyP := multilineP.Copy()
			replacer := strings.NewReplacer(replPairs...)
			walkTree(copyP, func(e *etree.Element) {
				if e.Tag == "t" {
					e.SetText(replacer.Replace(e.Text()))
				}
			})
			multilineParent.InsertChildAt(multilinePos, copyP)
			multilinePos += 1
		}
		multilineParent.RemoveChildAt(multilinePos)
	CONT:
		{
		}
	}
}
