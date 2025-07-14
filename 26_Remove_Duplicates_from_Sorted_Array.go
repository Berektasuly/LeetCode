import "fmt"

func removeDuplicates(nums []int) int {
	// var number []int
	// number = append(number, nums[0])
	// value := 0
	// for i := 0; i < len(nums); i++ {
	//     for j := 1; j < len(nums); j++ {
	//         if nums[i] == nums[j] {
	//             // number = append(number, nums[i])
	//             value += 1
	//         } else {
	//             number = append(number, nums[i])
	//         }
	//     }
	//     return value
	// }
	// return 0

	// number = append(number, nums[0])
	// for i := 0; i < len(nums); i++ {
	// 	for j := 0; j < len(number); j++ {
	// 		if number[j] == nums[i] {
	// 			value += 1
	// 		} else {
	// 			number = append(number, nums[i])
	// 		}
	// 	}
	// }
	// return value

	// var number []int
	// // number = append(number, nums[0])
	// value := 0
	// for i := 0; i < len(nums); i++ {
	// 	for j := 1; j < len(number); j++ {
    //         if nums[i] == nums[j] {
    //             value += 1
    //             fmt.Println("12313123")
    //         } else {
    //             number = append(number, nums[i])
    //             return number
    //         }
    //     }
    // }
    // fmt.Println(value)
    // return value

    var newarr []int // объявляем новый параметр
    newarr = append(newarr, nums[0]) // создаем массив с начальными значениями чтобы позже добавить новые не дубл-е
    number := nums[0] // Берет начальные значения для сравнения
    lut := 0 // счетчик для дубл-ых знач-й

    for i := 1; i < len(nums); i++ { //цикл для пробега по данному массиву
        if nums[i] != number { // если 2 != 1 то 
            number = nums[i] // число принимает новый вид
            newarr = append(newarr, nums[i]) // добавляет число сюда, не дубл
            fmt.Println(newarr)
        } else {
            lut += 1 // если дубл-но то +1
            // fmt.Println(lut)
        }
    }
    copy(nums, newarr)
    
    // fmt.Println(lut)
    return len(newarr)
}
