import XCTest
import SwiftTreeSitter
import TreeSitterQmljs

final class TreeSitterQmljsTests: XCTestCase {
    func testCanLoadGrammar() throws {
        let parser = Parser()
        let language = Language(language: tree_sitter_qmljs())
        XCTAssertNoThrow(try parser.setLanguage(language),
                         "Error loading Qmljs grammar")
    }
}
