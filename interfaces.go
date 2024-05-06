package main

import "fmt"

type Document interface {
	Doc() string
}

type EntityPerson struct {
	Name   string
	Age    uint8
	Status bool
}

func (p EntityPerson) String() string {
	return fmt.Sprintf("Name: %s, Age: %d", p.Name, p.Age)
}

type IndividualPerson struct {
	EntityPerson
	Lastname string
	cpf      string
}

func (p IndividualPerson) Doc() string {
	return fmt.Sprintf("my cpf: %s", p.cpf)
}

type LegalEntityPerson struct {
	EntityPerson
	CorporateName string
	cnpj          string
}

func (pj LegalEntityPerson) Doc() string {
	return fmt.Sprintf("my cnpj: %s", pj.cnpj)
}

func show(d Document) {
	//if d, ok := d.(IndividualPerson); ok {
	//	fmt.Println(d.Lastname)
	//}

	switch d := d.(type) {
	case IndividualPerson:
		fmt.Println("Individual:", d.Lastname)
	case LegalEntityPerson:
		fmt.Println("LegalEntity:", d.CorporateName)
	default:
		fmt.Println("Unknown document type", d)
	}

	fmt.Println(d)
	fmt.Println(d.Doc())
}

func main() {
	p := IndividualPerson{
		EntityPerson: EntityPerson{"John", 25, true},
		Lastname:     "Doe",
		cpf:          "123.456.789-00",
	}

	show(p)
	fmt.Println()

	pj := LegalEntityPerson{
		EntityPerson:  EntityPerson{"Wookye Ltda.", 10, true},
		CorporateName: "Wookye Tech",
		cnpj:          "123.456.789/1000-00",
	}

	show(pj)

}
