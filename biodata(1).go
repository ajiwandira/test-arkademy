package main
 
import (
    "fmt"
    "encoding/json"
)
 
type biodata struct {
	Name                string                      `json:"name"`
	Age                 int                         `json:"age"`
	Address             string                      `json:"address"`
	Hobbies             []string                    `json:"hobbies"`
	Is_married          bool                        `json:"is_married"`
    List_school         []map[string]interface{}    `json:"list_school"`
    Skills              []map[string]interface{}    `json:"skills"`
    Interest_in_coding  bool                        `json:"interest_in_coding"`
}
 
func main() {
    var resp biodata
    var list_school []map[string]interface{}
    var skills []map[string]interface{}

    resp.Name = "Aji Wandira"
    resp.Age = 19
    resp.Address = "JL.TANJUNG SANYANG"
    resp.Hobbies = []string{"climbing","games"}
    resp.Is_married = false

    smk := map[string]interface{}{
        "name": "SMKN 10 JAKARTA",
        "year_in": "2015",
        "year_out": "2018",
        "major": "RPL",
    }
    list_school = append(list_school, smk)

    smp := map[string]interface{}{
        "name": "SMPN 50 JAKARTA",
        "year_in": "2012",
        "year_out": "2015",
        "major": "",
    }
    list_school = append(list_school, smp)
    resp.List_school = list_school
    
    skill1 := map[string]interface{}{
        "skill_name": "C++",
        "level" : "beginner",
    }
    skills = append(skills, skill1)

    skill2 := map[string]interface{}{
        "skill_name": "C#",
        "level" : "beginner",
    }
    skills = append(skills, skill2)

    skill3 := map[string]interface{}{
        "skill_name": "Golang",
        "level" : "beginner",
    }
    skills = append(skills, skill3)
    resp.Skills = skills
    resp.Interest_in_coding = true

    strResponse, _ := json.Marshal(resp)
    fmt.Println(string(strResponse))

}