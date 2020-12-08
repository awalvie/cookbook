# written while reading the article on static code analysis
# at https://deepsource.io/blog/introduction-static-code-analysis/

import sys
import tokenize


class DoubleQuotesChecker:
    msg = "single quotes detected, use double quotes instead"
    def __init__(self):
        self.violations = []

    def find_violations(self, filename, tokens):
        for token_type, token, (line, col), _, _ in tokens:
            if (
                token_type == tokenize.STRING
                and (
                    token.startswith("'''")
                    or token.startswith("'")
                )
            ):
                self.violations.append((filename, line, col))

    def check(self, files):
        for filename in files:
            with tokenize.open(filename) as fd:
                tokens = tokenize.generate_tokens(fd.readline)
                self.find_violations(filename, tokens)

    def report(self):
        for violation in self.violations:
            filename, line, col = violation
            print(f"{filename}:{line}:{col}: {self.msg}")


if __name__ == '__main__':
    files = sys.argv[1:]
    checker = DoubleQuotesChecker()
    checker.check(files)
    checker.report()
