parser grammar Golearn;
import GolearnLexer;

expr: Atom | Atom OpBin Atom | OpUnary Atom | '(' expr ')';

arg: Type Id | Type Id ',';

statement: expr ';';

block: SymBlockL statement* SymBlockR;

fn: KeyFn Id block;
