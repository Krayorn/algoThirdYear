package main

import "fmt"

type Pers struct {
    Name string
    Ancestor string
    Age int
    isMale bool
    Children []Pers
}

type  childCount struct {
    Count int
    Pers Pers
}

func main(){
    var arr = make([]Pers, 10)

    arr[0] = Pers{
        Name: "Amélie",
        Age: 104,
        Ancestor: "",
        isMale: false,
    }

    arr[1] = Pers{
        Name: "Jack",
        Age: 86,
        Ancestor: "Amélie",
        isMale: true,
    }

    arr[2] = Pers{
        Name: "Pascal",
        Age: 59,
        Ancestor: "Jack",
        isMale: true,
    }

    arr[3] = Pers{
        Name: "Julien",
        Age: 33,
        Ancestor: "Pascal",
        isMale: true,
    }

    arr[4] = Pers{
        Name: "Caroline",
        Age: 30,
        Ancestor: "Pascal",
        isMale: false,
    }

    arr[5] = Pers{
        Name: "Olivier",
        Age: 59,
        Ancestor: "Jack",
        isMale: true,
    }

    arr[6] = Pers{
        Name: "Angélique",
        Age: 18,
        Ancestor: "Olivier",
        isMale: false,
    }

    arr[7] = Pers{
        Name: "Charlotte",
        Age: 22,
        Ancestor: "Olivier",
        isMale: false,
    }

    arr[8] = Pers{
        Name: "Lucien",
        Age: 5,
        Ancestor: "Olivier",
        isMale: true,
    }

    arr[9] = Pers{
        Name: "Roseline",
        Age: 79,
        Ancestor: "Amélie",
        isMale: false,
    }

    youngestAge := arr[0].Age
    youngestPers := arr[0]
    oldestAge := arr[0].Age
    oldestPers := arr[0]
    numberOfWomen := 0

    for _, v := range arr {
        if v.Age < youngestAge {
            youngestPers, youngestAge = v, v.Age
        } else if v.Age > oldestAge {
            oldestPers, oldestAge = v, v.Age
        }

        if !v.isMale {
            numberOfWomen++
        }
    }

    fmt.Println("youngest =", youngestPers)
    fmt.Println("oldest =", oldestPers)
    fmt.Println("numberOfWomen =", numberOfWomen)


    var tree = formatArray(arr)

    fmt.Println(getPersWithMostChildren(tree))

    newBorn := Pers{
        Name: "Barnabé",
        Age: 1,
        Ancestor: "Julien",
        isMale: true,
    }

    arr = append(arr, newBorn)

    getAscendance(newBorn, arr)
}

func formatArray(arr []Pers) Pers {
    var matriarch Pers
    for _, v := range arr {
        if v.Ancestor == "" {
            matriarch = v
            break
        }
    }

    matriarch.Children = getChildren(arr, matriarch)

    return matriarch
}

func getChildren(arr []Pers, parent Pers) []Pers {
    var children []Pers

    for _, v := range arr {
        if v.Ancestor == parent.Name {
            v.Children = getChildren(arr, v)
            children = append(children, v)
        }
    }

    return children
}

func getPersWithMostChildren(currentPers Pers) Pers {
    maxPers, currentMax := currentPers, len(currentPers.Children)
    for _, v := range currentPers.Children {
        a := getPersWithMostChildren(v)
        if len(a.Children) > currentMax {
            currentMax, maxPers = len(a.Children), a
        }
    }

    return maxPers
}

func getAscendance(origin Pers, allAncestor []Pers) {
    ancestorName := origin.Ancestor

    for _, v := range allAncestor {
        if v.Name == ancestorName {
            getAscendance(v, allAncestor)
            break
        }
    }

    fmt.Println(origin.Name)
}
