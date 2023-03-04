package main

import (
	"encoding/csv"
	"encoding/xml"
	"io"
	"os"
	"strconv"
	"strings"
)

// начало решения

// Employee представляет информацию о сотруднике
type Employee struct {
	ID       string `xml:"id,attr"`
	Name     string `xml:"name"`
	City     string `xml:"city"`
	Salary   int    `xml:"salary"`
	DeptCode string
}

// Organization представляет информацию об организации
type Organization struct {
	Depts []Department `xml:"department"`
}

// Department представляет информацию о департаменте
type Department struct {
	Code      string     `xml:"code"`
	Employees []Employee `xml:"employees>employee"`
}

// ConvertEmployees преобразует XML-документ с информацией об организации
// в плоский CSV-документ с информацией о сотрудниках
func ConvertEmployees(outCSV io.Writer, inXML io.Reader) error {
	// Декодируем XML-документ в структуру Organization
	org := Organization{}
	err := xml.NewDecoder(inXML).Decode(&org)
	if err != nil {
		return err
	}

	// Создаем объект csv.Writer для записи CSV-документа в outCSV
	csvWriter := csv.NewWriter(outCSV)

	// Записываем заголовок CSV-документа
	header := []string{"id", "name", "city", "department", "salary"}
	err = csvWriter.Write(header)
	if err != nil {
		return err
	}

	// Обходим каждый департамент и его сотрудников, записывая их в CSV-документ
	if len(org.Depts) == 0 {
		// Если нет департаментов, но есть заголовок, то записываем только заголовок

		csvWriter.Flush()
		if err := csvWriter.Error(); err != nil {
			return err
		}

		return nil
	}

	for _, dept := range org.Depts {
		for _, emp := range dept.Employees {
			emp.DeptCode = dept.Code
			record := []string{
				emp.ID,
				emp.Name,
				emp.City,
				emp.DeptCode,
				strconv.Itoa(emp.Salary),
			}
			err = csvWriter.Write(record)
			if err != nil {
				return err
			}
		}
	}

	csvWriter.Flush()
	if err := csvWriter.Error(); err != nil {
		return err
	}

	return nil
}

// конец решения

func main() {
	src := `<organization>
    <department>
        <code>hr</code>
        <employees>
            <employee id="11">
                <name>Дарья</name>
                <city>Самара</city>
                <salary>70</salary>
            </employee>
            <employee id="12">
                <name>Борис</name>
                <city>Самара</city>
                <salary>78</salary>
            </employee>
        </employees>
    </department>
    <department>
        <code>it</code>
        <employees>
            <employee id="21">
                <name>Елена</name>
                <city>Самара</city>
                <salary>84</salary>
            </employee>
        </employees>
    </department>
	<department>
        <code>ux</code>
        <employees>
            <employee id="20">
                <name>Лена</name>
                <city>Атырау</city>
            </employee>
        </employees>
		<employees>
            <employee id="200">
                <name>Эля</name>
                <city>Атырау</city>
				<salary>840</salary>
            </employee>
        </employees>
    </department>
	<department>
		<employees>
			<employee id="222">
                <name>Лена</name>
                <city>Атырау</city>
            </employee>
		</employees>
	</department>
</organization>`

	in := strings.NewReader(src)
	out := os.Stdout
	_ = ConvertEmployees(out, in)
	/*
		id,name,city,department,salary
		11,Дарья,Самара,hr,70
		12,Борис,Самара,hr,78
		21,Елена,Самара,it,84
	*/
}
