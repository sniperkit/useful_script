struct HasDrop;

impl Drop for HasDrop {
    fn drop(&mut self) {
        println!("Dropping");
    }
}

fn firework () {
    struct Firework {
        strength : i32,
    }

    impl Drop for Firework {
        fn drop(&mut self) {
            println!("BOOM times {}!!!!", self.strength);
        }
    }

    let fircracker = Firework{strength: 1};
    let tnt = Firework{ strength: 1000};
}

fn main() {
    let x = HasDrop;

    firework();
}
