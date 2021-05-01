/**
 * Definition for Employee.
 * type Employee struct {
 *     Id int
 *     Importance int
 *     Subordinates []int
 * }
 */

func calc(e *Employee, m map[int]*Employee) int {
    res := e.Importance
    for _, ee := range e.Subordinates {
        res += calc(m[ee], m)
    }
    return res
}


func getImportance(employees []*Employee, id int) int {
    m := make(map[int]*Employee)
    for _, e := range employees {
        m[e.Id] = e
    }
    return calc(m[id], m)
}
