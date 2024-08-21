//
//  SString.swift
//  smog
//
//  Created by Kristofer Younger on 8/17/23.
//

import Foundation

class SString: SObject {
    var s: String = ""
    
    init(s: String) {
        self.s = s
        super.init(nArgs: 0, clazz: Universe.shared.stringClass)
    }
    func string() -> String {
        return self.s
    }

    override func debugString() -> String {
        return "SString(\(String(describing: self.somClass())))"
    }
    
}
