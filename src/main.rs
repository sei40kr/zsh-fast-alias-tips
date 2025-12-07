mod lexer;
mod matcher;
mod model;
mod parser;

use std::io::{self, BufRead};

use lexer::Lexer;
use matcher::find_best_match;
use parser::Parser;

/// Reads alias definitions from stdin and finds the best match for the command.
///
/// # Arguments
///
/// * Command to match (provided as first command-line argument)
///
/// # Input
///
/// Reads alias definitions from stdin, one per line in the format:
/// - `name=expansion`
/// - `name='expansion'`
/// - `'name'='expansion'`
///
/// # Output
///
/// Prints the suggested alias to stdout:
/// - Full match: alias name only (e.g., `gst`)
/// - Partial match: alias name + remaining arguments (e.g., `gco feature-branch`)
fn main() {
    let args: Vec<String> = std::env::args().collect();

    if args.len() != 2 {
        eprintln!("Invalid number of arguments");
        std::process::exit(1);
    }

    let command = &args[1];

    let definitions = read_alias_definitions_from_stdin();

    if let Some(result) = find_best_match(&definitions, command) {
        let output = format_match_result(result, command);
        println!("{}", output);
    }
}

/// Reads and parses alias definitions from stdin.
fn read_alias_definitions_from_stdin() -> Vec<model::AliasDefinition> {
    let stdin = io::stdin();
    let reader = io::BufReader::with_capacity(1024, stdin.lock());

    reader
        .lines()
        .map_while(Result::ok)
        .filter_map(|line| {
            let mut lexer = Lexer::new(&line);
            let tokens = lexer.tokenize();
            let mut parser = Parser::new(tokens);
            parser.parse().ok()
        })
        .collect()
}

/// Formats the match result for output.
fn format_match_result(result: matcher::MatchResult, command: &str) -> String {
    if result.is_full_match {
        result.definition.name.clone()
    } else {
        format!(
            "{}{}",
            result.definition.name,
            &command[result.definition.expansion.len()..]
        )
    }
}
