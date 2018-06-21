extern crate iron;

use iron::prelude::*;
use iron::status;

fn main() {
    fn jabex(_: &mut Request) -> IronResult<Response> {
        Ok(Response::with((status::Ok, "Olá mundo! Escutem o SudoCast! :)")))
    }

    println!("On 7777");
    Iron::new(jabex).http("0.0.0.0:7777").unwrap();
}