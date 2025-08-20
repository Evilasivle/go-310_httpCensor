package main

import "testing"

func TestCatchHttp(t *testing.T) {
	testCases := [5]string{}
	testCases[0] = "Сообщение от дяди Васи, вы забыли миллион долларов: http://millionare.net"
	testCases[1] = "http://millionhere.com just take it!"
	testCases[2] = "http://spammysite.ru"
	testCases[3] = "Своё царство передаю вам http://newkingdom.com просто перейди по ссылке!"
	testCases[4] = "Следуй за белым кроликом: http://whiterabbit.ru"

	rightOutput := [5]string{
		"Сообщение от дяди Васи, вы забыли миллион долларов: http://**************",
		"http://*************** just take it!",
		"http://*************",
		"Своё царство передаю вам http://************** просто перейди по ссылке!",
		"Следуй за белым кроликом: http://**************",
	}

	for index, val := range testCases {
		result := CatchHttp(val)
		if result != rightOutput[index] {
			t.Errorf("CatchHttp %s, ожидалось %s, получилось %s", val, rightOutput[index], result)
		}
	}
}
