//macro_rules声明一个宏，say_hello是宏名字，say_hello后边是一对大括号，大括号中间是宏的内容。
//宏的内容中=>前边的()表示这个宏没有输入参数。在整个宏的内部，我们调用的println宏。
macro_rules! say_hello {
    () => (
        println!("Hello");
        )
}

//宏的内容除了可以放在大括号中以外还可以放在()和[]中。
//这里要注意的是当宏内容放在大括号中时，末尾不用放“;”，而小括号和中括号后要放";".

macro_rules! say_hello1 (
    () => ( 
        println!("Hello one");
        )
 );

macro_rules! say_hello2 [
    () => (
        println!("Hello two");
        )
];

//宏的参数是包含在宏的内容中的。
//宏定义中宏名称后边大括号中的内容称为macro rule(宏规则)。一个宏可以包含任意多个宏规则。
//宏规则的格式是这样的:
//  () => {};
//  () => [];
//  () => ();
//每个宏规则有一个=>。
//  在=>左边的小括号是该规则的输入参数，该参数的形式是一个叫做matcher的匹配器。
//  在=>右边的的括号是该规则对应的宏内容，需要执行或者被替换，叫做transcriber(转换器)，可以是语句，表达式等。
//
//within a matcher, each metavariable has a fragment specifier, identifying which syntatic for it
//matches.
//  1. ident: an identifier. Examples: x; foo.
//  2. path:  a qualified name. Example: T::SpecialA.
//  3. expr:  an expresssion. Examples: 2 + 2; if true {1} else {0};.
//  4. ty:    a type. Examples: i32; Vec<(char, String)>; &T,
//  5. pat:   a pattern. Examples: Some(t); (17, 'a');_.
//  6. stmt:  a single statement. Example: let x = 3.
//  7. block: a brace-delimited sequence of statements and optionally an expression. Example:
//          { log(error, "hi"); return 12;}.
//  8. item:  an item. Exmaples: fn foo() {}; struct Bar;
//  9. meta:  a 'meta item", as found in attributes. Example: cfg(target_os="windows").
//  10. tt:   a single token tree.
macro_rules! print_hello {
    ($msg:expr) =>  (
        println!("Hello {}", $msg);
        )
}

//宏的输入不是通过参数列表传入的，而是通过match的方式进行查找或者匹配的。
macro_rules! foo {
    (x => $e:expr) => (println!("Mode X: {}", $e));
    (y => $e:expr) => (println!("Mode Y: {}", $e));
}

fn main() {
    say_hello!();
    say_hello1!();
    say_hello2!();
    print_hello!("liwei");

    foo!(x => 7);
    foo!(y => 3);
}


