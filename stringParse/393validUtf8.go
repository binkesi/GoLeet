package stringparse

// https://leetcode-cn.com/problems/utf-8-validation/

func validUtf8(data []int) bool {
	//‘1’的个数
	nums1 := 0
	lenght := 0
	for lenght < len(data) {
		//nums1>0 判断字节首两位是否为‘10’
		if nums1 > 0 {
			if !(data[lenght] < 192 && data[lenght] >= 128) {
				return false
			}
			nums1--
		} else {
			//nums1=0
			//1、如果首位为 0 直接跳过
			//2、如果字节首两位是为‘10’则返回false
			//3、开始判断有几个1在字节首位 、重新给nums1赋值,
			//如果首位为 0 直接跳过
			computenum := 128
			num := data[lenght]
			//如果num比128小，那么就是说是1字节的
			if num < computenum {
				lenght++
				continue
			}
			//如果字节首两位是为‘10’则返回false
			if num < 192 && num >= computenum {
				return false
			}
			for ; computenum > 0; computenum /= 2 {
				if num >= computenum {
					nums1++
					//‘1’的个数必须要在5个以内
					if nums1 > 4 {
						return false
					}
					num -= computenum
				} else {
					//否则直接跳出这个循环
					nums1--
					break
				}
			}
		}
		lenght++
	}
	if nums1 > 0 {
		return false
	}
	return true
}
