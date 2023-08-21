//
//  SSymbol.swift
//  smog
//
//  Created by Kristofer Younger on 8/17/23.
//

import Foundation

class SSymbol: SString {
    var numSignatureArguments: Int = 0
    
    override init(s: String) {
        super.init(s: s)
        self.determineNumberOfArguments()
    }
    
    func determineNumberOfArguments() {
        // Check for binary signature
        if self.isBinarySignature() {
            self.numSignatureArguments = 2
        } else {
            // Count the colons in the signature string
            var numberOfColons = 0

            // Iterate through every character in the signature string
            for c in self.s {
                if c == ":" {
                    numberOfColons += 1
                }
            }
            // The number of arguments is equal to the number of colons plus one
            self.numSignatureArguments = numberOfColons + 1
        }
    }
    
    func isBinarySignature() -> Bool {
        for c in self.s {
            if c != "~" && c != "&" && c != "|" && c != "*" &&
                c != "/" && c != "@" && c != "+" && c != "-" &&
                c != "=" && c != ">" && c != "<" && c != "," &&
                c != "%" && c != "\\" {
                return false
            }
        }
        return true
    }

    
    override func debugString() -> String {
        return "SSymbol(\(String(describing: self.somClass())))"
    }

}
