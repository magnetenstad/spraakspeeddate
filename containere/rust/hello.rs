use std::collections::HashMap;

fn main() {
    println!("Hello World!");
    println!("{}", find_unique(&[1, 2, 3, 2, 1])); // expect 3
    println!("{}", find_unique(&[5000])); // expect 5000
    println!("{}", find_unique(&[1, 1, 1, 1, 1, 5, 9, 1, 9])); // expect 5
}

fn find_unique(array: &[i32]) -> i32 {
    let mut count = HashMap::new();
    for num in array.iter() {
        count.insert(num, 1 + 
            match count.get(num) {
                Some(x) => x,
                None    => &0,
            });
    }
    for (key, value) in count.into_iter() {
        if value == 1 {
            return *key;
        }
    }
    return 0
}
