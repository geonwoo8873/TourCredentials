// `std`는 Rust에서 기본 제공되는 표준 라이브러리
use std::io;
use std::cmp::Ordering;
use rand::prelude::*;

fn main() {
    // Update : `rand`의 라이브러리는 "0.9.0" 버전에서 `thread_rng`와 `gen_range`에서 `random_range`로 변경되었다. 
    // rand lib URL : https://rust-random.github.io/book/update-0.9.html
    // `random_range(start..end)`가 스레드[Thread]에 저장할 값을 임의의 숫자로 생성한다.
    let random_value = rand::rng().random_range(1..=100);
    let random_data = rand::rng().random::<u128>();

    let guess = loop {
        // `String::new(string)`는 값을 받았을 때 값을 `String`타입으로 변수[Instance] 생성
        let mut user_input = String::new();

        // `io::stdin`는 사용자로 부터 입력을 요청하여 `read_line(&mut value)`에 값을 읽을 수 있게 한다.
        // `expect(string)`은 요청 값이 타입에 일치하지 않다면 예외 경고 결과 출력을 반환한다.
        io::stdin().read_line(&mut user_input).expect("Failed to read line");

        // `trim()`은 문자열의 앞과 뒤에 있는 공백을 제거한다. `parse()`는 문자열을 숫자로 변환한다.
        // let guess: u32 = user_input.trim().parse().expect("Failed to read line");
        let guess: u32 = match user_input.trim().parse() {
            // `Ok`는 `Result` 타입의 열거형으로, 성공적으로 값을 반환할 때 사용된다.
            Ok(num) => num,
            // `Err`는 `Result` 타입의 열거형으로, 오류가 발생했을 때 사용된다.
            Err(_) => {
                println!("Please enter a valid number!");
                continue;
            }
        };

        // `match`는 `arm[갈래]`들로 이루어져 하나의 `패턴[pattern]`과 `match` 표현식에서 주어진 값이 패턴과 일치하면 실행하는 구조이다.
        match guess.cmp(&random_value) {
            // `user value < random value`
            Ordering::Less => println!("Up"),
            // `user value > random value`
            Ordering::Greater => println!("Down"),
            // `user value == random value`
            Ordering::Equal => {
                println!("Equal");
                break guess;
            }
        };
    };

    println!("Random Number : {random_value} / {random_data} , Get Guess : {guess}");
}