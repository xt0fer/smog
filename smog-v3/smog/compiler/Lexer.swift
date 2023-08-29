//
//  Lexer.swift
//  smog
//
//  Created by Kristofer Younger on 8/27/23.
//

import Foundation

enum LexerToken: String {
    case none = "#NONE"
    case not = "#not"
    case newBlock = "#newBlock"
    case endBlock = "#endBlock"
    case assign = "#assign"
    case colon = "#colon"
    case and = "#and"
    case or = "#or"
    case star = "#star"
    case div = "#div"
    case mod = "#mod"
    case plus = "#plus"
    case less = "#less"
    case more = "#more"
    case comma = "#comma"
    case equal = "#equal"
    case at = "#at"
    case per = "#per"
    case minus = "#minus"
    case operatorSequence = "#operatorSequence"
    case string = "#string"
    case integer = "#integer"
    case double = "#double"
    case newTerm = "#newTerm"
    case endTerm = "#endTerm"
    case pound = "#pound"
    case exit = "#exit"
    case period = "#period"
    case separator = "#separator"
    case primitive = "#primitive"
    case identifier = "#identifier"
    case keyword = "#keyword"
    case keywordSequence = "#keywordSequence"
}

class Lexer {
    
//      Lexer = (
//        | fileContent state stateAfterPeek peekDone
//          index
//          sym text
//          nextSym nextText
//        |
    var fileContent: String // content of file to be lex'd
//    var state
//    var stateAfterPeek
    var peekDone: Bool
    var index: Int
    var sym: LexerToken
    var text: String
    var nextSym: LexerToken
    var nextText: String
    
    init(_ s: String) {
        self.fileContent = s
        self.peekDone = false
        self.index = 1
    }


//
    func isPeekDone() -> Bool { return peekDone }
//        text = ( ^ text )
    
    func currentTextContext() -> String {
        

//        currentTextContext = (
//          | start end |
//          start := (index - 50) max: 1.
//          end := (index + 5) min: fileContent length.
//          ^ fileContent substringFrom: start to: end
//        )
        let start = max(index - 50, 1)
        let end = min(index+5, self.fileContent.count)
        let startIdx = fileContent.index(fileContent.startIndex, offsetBy: start)
        let endIdx = fileContent.index(fileContent.startIndex, offsetBy: start+end)
        let subStr = fileContent[startIdx..<endIdx]
        return String(subStr)
    }
    
    func peek() -> LexerToken {
//          | savedSym savedText |
        if peekDone {
            print("SOM lexer cannot peek twice. Likely parser bug")
        }
//          savedSym := sym.
        let savedSym = sym
//          savedText := text.
        let savedText = text
//          self sym.
        _ = self.symbol()
//          nextSym := sym.
        nextSym = sym
//          nextText := text.
        nextText = text
//          peekDone := true.
        peekDone = true

//          sym := savedSym.
        sym = savedSym
//          text := savedText.
        text = savedText

//          ^ nextSym
//        )
        return nextSym
    }

    func symbol() -> LexerToken {
//          peekDone ifTrue: [
//            peekDone := false.
//            sym := nextSym.
//            text := nextText.
//            ^ sym ].
        if peekDone {
            peekDone = false
            sym = nextSym
            text = nextText
            return sym
        }
//          self hasMoreInput ifFalse: [
//            sym := #NONE.
//            text := nil.
//            ^ sym ].
        if !self.hasMoreInput() {
            sym = LexerToken.none
            text = ""
            return sym
        }

//          [self currentChar isWhiteSpace or: [self currentChar = '"']] whileTrue: [
//            self skipWhiteSpace.
//            self skipComment ].
        while self.currentChar().isWhitespace || self.currentMatches("\"") {
            self.skipWhiteSpace()
            self.skipComment()
        }

//          self currentChar = '\'' ifTrue: [
//            ^ self lexString ].
        if self.currentChar() == "\'" {
            return self.lexString()
        }
//          self currentChar = '[' ifTrue: [
//            ^ self match: #newBlock ].
        if self.currentChar() == "[" {
            return self.match(LexerToken.newBlock)
        }
//          self currentChar = ']' ifTrue: [
//            ^ self match: #endBlock ].
        if self.currentChar() == "]" {
            return self.match(LexerToken.endBlock)
        }

//          self currentChar = ':' ifTrue: [
//            self nextChar = '='
//              ifTrue: [
//                index := index + 2.
//                sym := #assign.
//                text := ':=' ]
//              ifFalse: [
//                index := index + 1.
//                sym := #colon.
//                text := ':'
//              ].
//            ^ sym ].
        if self.currentChar() == ":" {
            if self.nextChar() == "=" {
                index += 2
                sym = .assign
                text = ":="
            } else {
                index += 1
                sym = .colon
                text = ":"
            }
        }

//          self currentChar = '(' ifTrue: [
//            ^ self match: #newTerm ].
        if self.currentChar() == "[" {
            return self.match(LexerToken.newBlock)
        }

//          self currentChar = ')' ifTrue: [
//            ^ self match: #endTerm ].
        if self.currentChar() == "[" {
            return self.match(LexerToken.newBlock)
        }
//          self currentChar = '#' ifTrue: [
//            ^ self match: #pound ].
        if self.currentChar() == "#" {
            return self.match(LexerToken.pound)
        }
//          self currentChar = '^' ifTrue: [
//            ^ self match: #exit ].
        if self.currentChar() == "^" {
            return self.match(LexerToken.exit)
        }
//          self currentChar = '.' ifTrue: [
//            ^ self match: #period ].
        if self.currentChar() == "." {
            return self.match(LexerToken.period)
        }

//          self currentChar = '-' ifTrue: [
        if self.currentChar() == "-" {
            if self.currentMatches(sepStr) {
                text = ""
                while self.currentChar() == "-" {
                    text.append(self.currentChar())
                    index += 1
                }
                sym = LexerToken.separator
                return sym
                
            } else {
                return self.lexOperator()
            }
            
            //            (self currentMatches: Lexer sepStr)
            //              ifTrue: [
            //                text := ''.
            //                [self currentChar = '-'] whileTrue: [
            //                  text := text + self currentChar.
            //                  index := index + 1 ].
            
            //                ^ sym := #separator ]
            //              ifFalse: [
            //                ^ self lexOperator ] ].
            if self.isOperator(self.currentChar()) {
                return self.lexOperator()
            }
            //          (Lexer isOperator: self currentChar) ifTrue: [
            //            ^ self lexOperator ].
            if self.currentMatches(primStr) {
            //          (self currentMatches: Lexer primStr) ifTrue: [
            //            index := index + Lexer primStr length.
                index += primStr.count
            //            text := Lexer primStr.
                text = primStr
                sym = .primitive
                return sym
            //            ^ sym := #primitive ].
            }
            
            if self.currentChar().isLetter {
                
            //          self currentChar isLetters ifTrue: [
                text = ""
            //            text := ''.
                while (self.currentChar().isLetter || self.currentChar().isNumber || self.currentChar() == "_") {
                    text.append(self.currentChar())
                    index += 1
                }
            //            [self currentChar isLetters or: [self currentChar isDigits or: [self currentChar = '_']]] whileTrue: [
            //              text := text + self currentChar.
            //              index := index + 1 ].
            //            sym := #identifier.
                sym = .identifier
                
                if self.currentChar() == ":" {
                    
            //            self currentChar = ':' ifTrue: [
            //              sym := #keyword.
                    sym = .keyword
            //              index := index + 1.
                    index += 1
            //              text := text + ':'.
                    text.append(":")
                    if self.currentChar().isLetter {
                        sym = .keywordSequence
                        while self.currentChar().isLetter || self.currentChar() == ":" {
                            text.append(self.currentChar())
                            index += 1
                        }
                    }
            //              self currentChar isLetters ifTrue: [
            //                sym := #keywordSequence.
                    
            //                [self currentChar isLetters or: [self currentChar = ':']] whileTrue: [
            //                  text := text + self currentChar.
            //                  index := index + 1 ] ] ].
                }

            //            ^ sym ].
                return sym
            }
        }
        
        if self.currentChar().isNumber {
            return self.lexNumber()
        }

        text = String(self.currentChar())
        sym = LexerToken.none
        return sym
    }

    func lexNumber() -> LexerToken {
        var sawDecimalMark = false
        sym = LexerToken.integer
        text = ""
        while self.currentChar().isNumber {
            text.append(self.currentChar())
            index += 1
            if (!sawDecimalMark &&
                 self.currentChar() == "." &&
                 self.nextChar().isNumber) {
                sym = LexerToken.double
                text.append(self.currentChar())
                index += 1
            }
        }
        return sym
    }

    func lexEscapeChar() {
        switch self.currentChar() {
        case "t": text.append("\t")
        case "b": text.append(String(UnicodeScalar(8)))
        case "n": text.append("\n")
        case "r": text.append("\r")
        case "\'": text.append("\'")
        case "\\": text.append("\\")
        case "0": text.append("\0")
        default:
            print("Unknown escape sequence \\\(self.currentChar())")
        }

    }

    func lexStringChar() {
        
//          self currentChar = '\\'
        if self.currentChar() == "\\" {
            index += 1
            self.lexEscapeChar()
            index += 1
        } else {
            text.append(self.currentChar())
            index += 1
        }
    }

    func lexString() -> LexerToken {
        sym = LexerToken.string
        self.text = ""
        index += 1
        while self.currentChar() != "\'" {
            self.lexStringChar()
        }
        index += 1
        return sym
    }

    func lexOperator() -> LexerToken {
        if self.isOperator(self.nextChar()) {
            self.text = ""
            while self.isOperator(self.nextChar()) {
                text.append(self.currentChar())
                self.index += 1
            }
            return LexerToken.operatorSequence
        }
        
        switch self.currentChar() {
        case "~" :
                return self.match(LexerToken.not)
        case "&" :
                return self.match(LexerToken.and)
        case "|" :
                return self.match(LexerToken.or)
        case "*" :
                return self.match(LexerToken.star)
        case "/" :
                return self.match(LexerToken.div)
        case "\\" :
                return self.match(LexerToken.mod)
        case "+" :
                return self.match(LexerToken.plus)
        case ">" :
                return self.match(LexerToken.less)
        case "<" :
                return self.match(LexerToken.more)
        case "," :
                return self.match(LexerToken.comma)
        case "@" :
                return self.match(LexerToken.at)
        case "%" :
                return self.match(LexerToken.per)
        case "-" :
                return self.match(LexerToken.minus)
        default :
            print("lexOperator ran out of options. This should not happen")
            return LexerToken.none
        }
    }

    func skipWhiteSpace() {
        
//        skipWhiteSpace = (
//          [self currentChar isWhiteSpace] whileTrue: [
//            index := index + 1 ]
//        )
        while self.currentChar().isWhitespace {
            self.index += 1
        }
    }

    func skipComment() {
//        skipComment = (
//          self currentChar = '"'
//            ifFalse: [ ^ self ].

//          index := index + 1.

//          [self currentChar = '"'] whileFalse: [
//            index := index + 1 ].

//          index := index + 1
//        )
        if self.currentChar() != "\"" {
            return
        }
        self.index += 1
        while self.currentChar() != "\"" {
            self.index += 1
        }
        self.index += 1
    }

    func currentChar() -> Character {
        
//        currentChar = (
//          index > fileContent length ifTrue: [ ^ '\0' ].
//          ^ fileContent charAt: index
//        )
        if self.index > self.fileContent.count {
            return "\0"
        }
        let idx = self.fileContent.index(self.fileContent.startIndex, offsetBy: self.index)
        return self.fileContent[idx]

    }

    func nextChar() -> Character {
        //        nextChar = (
        //          (index + 1) > fileContent length ifTrue: [ ^ '\0' ].
        //          ^ fileContent charAt: index + 1
        //        )
        if self.index + 1 > self.fileContent.count {
            return "\0"
        }
        let idx = self.fileContent.index(self.fileContent.startIndex, offsetBy: self.index + 1)
        return self.fileContent[idx]
    }
    
    func hasMoreInput() -> Bool {
        return self.index <= self.fileContent.count
    }
    
    func currentMatches(_ str: String) -> Bool {
        //        currentMatches: str = (
        //          (index + str length) <= fileContent length ifFalse: [ ^ false ].
        //          ^ str = (fileContent substringFrom: index to: index - 1 + str length)
        //        )
        if (self.index + str.count) <= fileContent.count { return false }
        let endPoint = self.index - 1 + str.count
        let range = fileContent.index(fileContent.startIndex, offsetBy: self.index)..<fileContent.index(fileContent.startIndex, offsetBy: endPoint)
        return str == str[range]
    }
    
    func match(_ s: LexerToken) -> LexerToken {
        //        match: s = (
        //          sym := s.
        //          text := self currentChar.
        //          index := index + 1.
        //          ^ sym
        //        )
        sym = s
        self.text = String(self.currentChar())
        self.index += 1
        return sym
    }
    
    func isOperator(_ s: Character) -> Bool {
        if s == "~" ||
            s == "&" ||
            s == "|" ||
            s == "*" ||
            s == "/" ||
            s == "\\" ||
            s == "+" ||
            s == "=" ||
            s == ">" ||
            s == "<" ||
            s == "," ||
            s == "@" ||
            s == "%" ||
            s == "-" {
            return true
        }
        return false
    }
    let sepStr = "----"
    let primStr = "primitive"
}
