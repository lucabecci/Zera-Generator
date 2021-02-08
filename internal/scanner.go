package internal

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	. "github.com/logrusorgru/aurora/v3"
)

type Scanner struct {
	reader bufio.Reader
}

func New() *Scanner {
	reader := bufio.NewReader(os.Stdin)
	scanner := Scanner{
		reader: *reader,
	}
	return &scanner
}

func (s *Scanner) Run() {
	project, err := s.ProjectName()
	if err != nil {
		log.Fatal(err)
	}
	s.SelectTemplate(project)
}

func (s *Scanner) ProjectName() (string, error) {
	fmt.Print(Magenta("Project Name:"))
	fmt.Print(Magenta("->"))
	//get user input
	text, _ := s.reader.ReadString('\n')
	text = strings.Replace(text, "\n", "", -1)
	//check if the user used the :exit asignation for close the cli
	if strings.Compare(":exit", text) == 0 {
		os.Exit(1)
	}
	//create project
	cmd := exec.Command("mkdir", text)
	_ = cmd.Run()
	return text, nil
}

func (s *Scanner) SelectTemplate(p string) {
	fmt.Print(Magenta("Template Name:"))
	fmt.Print(Magenta("->"))
	//creating a template manager
	command := TemplateGenerator(p)
	//get template
	text, _ := s.reader.ReadString('\n')
	text = strings.Replace(text, "\n", "", -1)
	//check if the user used the :exit asignation for close the cli
	if strings.Compare(":exit", text) == 0 {
		os.Exit(1)
	}
	//creating template
	command.Compare(text)
}
