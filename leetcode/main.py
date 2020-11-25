#!/bin/python3
import json
import sys
import os
import subprocess

PWD = os.path.dirname(os.path.abspath(__file__))

def color(num, *args, **kwargs):
    end = kwargs.get("end", "\n")
    if "end" in kwargs:
        del kwargs["end"]
    print("\033[%sm" % (str(num)), end="")
    print(*args, **kwargs, end="")
    print("\033[0m", end=end)


class Context:
    def __init__(self, _id, name, tests=[], language="python", debug=False, verbose=False):
        self.id = _id
        self.language = language
        self.debug = debug
        self.name = name
        self.tests = tests
        self.verbose = verbose

    def __repr__(self):
        return "id={} name={} language={} debug={} verbose={}\n{}".format(
            self.id, self.name, self.language,
            self.debug, self.verbose, self.tests
        )


def call(*args, **kwargs):
    proc = subprocess.Popen(
        " ".join(args),
        shell=True,
        stdout=subprocess.PIPE,
        stdin=subprocess.PIPE,
        bufsize=0
    )
    try:
        outs, errs = proc.communicate(
            timeout=15,
            input=bytes(
                json.dumps(kwargs["input"]),
                encoding="utf-8",
            ) if "input" in kwargs else None,
        )
        outs = str(outs, encoding="utf-8")
    except Exception as e:
        outs = ""
        color(1, e)
        color(1, errs)
    finally:
        proc.kill()
    try:
        res = json.loads(outs)["out"]
    except:
        res = outs
    return  res


class CodeRunner():
    def __init__(self, language, suffix, cmd, codeFormat):
        self.language = language
        self.suffix = suffix
        self.cmd = cmd
        self.codeFormat = codeFormat
        self.PWD = os.path.dirname(os.path.abspath(__file__))

    def read(self, ctx):
        folder = "%s/%s" % (self.PWD, ctx.id)
        source = os.path.join(folder, "%s.%s" % (ctx.id, self.suffix))

        with open(source) as f:
            code = f.read()
        return code

    def generate(self, ctx, code):
        folder = "%s/%s" % (self.PWD, ctx.id)
        tmpFolder = os.path.join(folder, "dist")
        tmpFile = os.path.join(tmpFolder, "%s.%s" % (ctx.id, self.suffix))

        if not os.path.exists(tmpFolder):
            os.mkdir(tmpFolder)
        with open(tmpFile, "w") as f:
            f.write(self.codeFormat.format(code=code, name=ctx.name))

        return tmpFile

    def afterGenerate(self, ctx, tmpFile):
        pass

    def check(self, ctx, tmpFile):
        total = len(tests)
        passed = True
        for idx, test in enumerate(ctx.tests):
            if ctx.verbose:
                color("1;34", "Test", idx+1)

            out = call(self.cmd, tmpFile, input=test["in"])

            if out != test["out"]:
                passed = False
                color("1;31", "Error at", idx + 1, "test")

            if out != test["out"] or ctx.verbose:
                color(3, "\tinput\t", end="")
                color(0, test["in"])

                color(3, "\tgot  \t", end="")
                color(0, out)

                color(3, "\twant \t", end="")
                color(0, test["out"])

        if passed:
            color(92, "All %d testcases passed" % total)

    def run(self, ctx):
        code = self.read(ctx)
        tmpFile = self.generate(ctx, code)
        self.afterGenerate(ctx, tmpFile)
        self.check(ctx, tmpFile)


py3fmt = '''
from typing import *
import json

{code}

if __name__ == "__main__":
    solution = Solution()
    print(json.dumps({{ 
        "out": solution.{name}(json.loads(input())) 
    }}), end="")
'''


class Python3Runner(CodeRunner):
    def __init__(self):
        super().__init__("python3", "py", "python3", py3fmt)


    if passed:
        print("\033[92mAll %d testcases passed\033[0m" % total)
'''


def runPython3(_id, name, tests):
    folder = "%s/%s" % (PWD, _id)
    source = os.path.join(folder, "%s.py" % _id)
    tmpFolder = os.path.join(folder, "dist")
    tmpFile = os.path.join(tmpFolder, "%s.py" % _id)
    with open(source) as f:
        code = f.read()
    if not os.path.exists(tmpFolder):
        os.mkdir(tmpFolder)
    with open(tmpFile, "w") as f:
        f.write(py3template.format(code=code, name=name, tests=tests))

    res = call("python3", tmpFile)
    print(res)


if __name__ == "__main__":
    _id = sys.argv[1]
    with open("./%s/problem.json" % _id) as f:
        problem = json.loads(f.read())
    name = problem["name"]
    tests = problem["samples"]
    runPython3(_id, name, tests)
