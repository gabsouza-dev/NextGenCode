use rand::Rng;
use chrono::Local;

fn main() {
    println!("Hello, World!");
  
    let num = rand::thread_rng().gen_range(1..101);
  
    println!("{}", num);
    println!("{}", Local::now());
}
