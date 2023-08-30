//
//  main.swift
//  smog
//
//  Created by Kristofer Younger on 8/17/23.
//

import Foundation

print("Hello, World!")


//Main = (
//  run: args = (
//    | u args2 |
//    u := Universe new.
//    args2 := args copyFrom: 2.
//    u interpret: args2.
//    u exit: 0.
//  )
//)

struct Main {
    
    func run(args: [String]) {
        let u = Universe()
        let args2 = args.suffix(1)
        
        u.interpret(args2)
        
        
        u.exit(0)
    }
}

Main().run(args: ["Hello"])

