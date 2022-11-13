# gettext-report

A utility to analyze GNU GetText PO files to report the amount text that has been translated.

## Installation

```
go install github.com/brianpursley/gettext-report@latest
```

## Usage

```
Usage:
  gettext-report [.po files...] [flags]

Examples:

# Generate a report for two specific .po files
gettext-report french.po german.po

# Generate a report for two specific .po files, including a template file
gettext-report french.po german.po -t template.pot

# (Bash) Generate a report for all .po files under the current directory, including a template file
gettext-report $(find . -name *.po) -t template.pot

Flags:
  -h, --help              help for gettext-report
  -o, --output string     Output format: table or json (default "table")
  -t, --template string   Template (.pot) file
```

## Examples

### All .po files in a directory

```
$ gettext-report $(find staging/src/k8s.io/kubectl -name *.po)
File                                                                                      Translated  Total  Percent
staging/src/k8s.io/kubectl/pkg/util/i18n/translations/kubectl/de_DE/LC_MESSAGES/k8s.po    92          92     100.0 %
staging/src/k8s.io/kubectl/pkg/util/i18n/translations/kubectl/default/LC_MESSAGES/k8s.po  0           336    0.0 %
staging/src/k8s.io/kubectl/pkg/util/i18n/translations/kubectl/en_US/LC_MESSAGES/k8s.po    0           336    0.0 %
staging/src/k8s.io/kubectl/pkg/util/i18n/translations/kubectl/fr_FR/LC_MESSAGES/k8s.po    7           7      100.0 %
staging/src/k8s.io/kubectl/pkg/util/i18n/translations/kubectl/it_IT/LC_MESSAGES/k8s.po    101         101    100.0 %
staging/src/k8s.io/kubectl/pkg/util/i18n/translations/kubectl/ja_JP/LC_MESSAGES/k8s.po    12          101    11.9 %
staging/src/k8s.io/kubectl/pkg/util/i18n/translations/kubectl/ko_KR/LC_MESSAGES/k8s.po    7           7      100.0 %
staging/src/k8s.io/kubectl/pkg/util/i18n/translations/kubectl/pt_BR/LC_MESSAGES/k8s.po    100         101    99.0 %
staging/src/k8s.io/kubectl/pkg/util/i18n/translations/kubectl/zh_CN/LC_MESSAGES/k8s.po    101         101    100.0 %
staging/src/k8s.io/kubectl/pkg/util/i18n/translations/kubectl/zh_TW/LC_MESSAGES/k8s.po    7           7      100.0 %
staging/src/k8s.io/kubectl/pkg/util/i18n/translations/test/default/LC_MESSAGES/k8s.po     2           2      100.0 %
staging/src/k8s.io/kubectl/pkg/util/i18n/translations/test/en_US/LC_MESSAGES/k8s.po       2           2      100.0 %
TOTAL                                                                                     431         1193   36.1 %
```

### All .po files in a directory, with a template (.pot) file

```
$ gettext-report $(find staging/src/k8s.io/kubectl -name *.po) -t staging/src/k8s.io/kubectl/pkg/util/i18n/translations/kubectl/template.pot
File                                                                                      Translated  Total  Percent
staging/src/k8s.io/kubectl/pkg/util/i18n/translations/kubectl/de_DE/LC_MESSAGES/k8s.po    92          336    27.4 %
staging/src/k8s.io/kubectl/pkg/util/i18n/translations/kubectl/default/LC_MESSAGES/k8s.po  0           336    0.0 %
staging/src/k8s.io/kubectl/pkg/util/i18n/translations/kubectl/en_US/LC_MESSAGES/k8s.po    0           336    0.0 %
staging/src/k8s.io/kubectl/pkg/util/i18n/translations/kubectl/fr_FR/LC_MESSAGES/k8s.po    7           336    2.1 %
staging/src/k8s.io/kubectl/pkg/util/i18n/translations/kubectl/it_IT/LC_MESSAGES/k8s.po    101         336    30.1 %
staging/src/k8s.io/kubectl/pkg/util/i18n/translations/kubectl/ja_JP/LC_MESSAGES/k8s.po    12          336    3.6 %
staging/src/k8s.io/kubectl/pkg/util/i18n/translations/kubectl/ko_KR/LC_MESSAGES/k8s.po    7           336    2.1 %
staging/src/k8s.io/kubectl/pkg/util/i18n/translations/kubectl/pt_BR/LC_MESSAGES/k8s.po    100         336    29.8 %
staging/src/k8s.io/kubectl/pkg/util/i18n/translations/kubectl/zh_CN/LC_MESSAGES/k8s.po    101         336    30.1 %
staging/src/k8s.io/kubectl/pkg/util/i18n/translations/kubectl/zh_TW/LC_MESSAGES/k8s.po    7           336    2.1 %
staging/src/k8s.io/kubectl/pkg/util/i18n/translations/test/default/LC_MESSAGES/k8s.po     2           336    0.6 %
staging/src/k8s.io/kubectl/pkg/util/i18n/translations/test/en_US/LC_MESSAGES/k8s.po       2           336    0.6 %
TOTAL                                                                                     431         4032   10.7 %
```

### JSON output
```
$ gettext-report $(find staging/src/k8s.io/kubectl -name *.po) -t staging/src/k8s.io/kubectl/pkg/util/i18n/translations/kubectl/template.pot -o json
{
  "Count": 4032,
  "Diff": 431,
  "Percent": 10.689485,
  "Files": [
    {
      "File": "staging/src/k8s.io/kubectl/pkg/util/i18n/translations/kubectl/de_DE/LC_MESSAGES/k8s.po",
      "Count": 336,
      "Diff": 92,
      "Percent": 27.380953
    },
    {
      "File": "staging/src/k8s.io/kubectl/pkg/util/i18n/translations/kubectl/default/LC_MESSAGES/k8s.po",
      "Count": 336,
      "Diff": 0,
      "Percent": 0
    },
    {
      "File": "staging/src/k8s.io/kubectl/pkg/util/i18n/translations/kubectl/en_US/LC_MESSAGES/k8s.po",
      "Count": 336,
      "Diff": 0,
      "Percent": 0
    },
    {
      "File": "staging/src/k8s.io/kubectl/pkg/util/i18n/translations/kubectl/fr_FR/LC_MESSAGES/k8s.po",
      "Count": 336,
      "Diff": 7,
      "Percent": 2.0833335
    },
    {
      "File": "staging/src/k8s.io/kubectl/pkg/util/i18n/translations/kubectl/it_IT/LC_MESSAGES/k8s.po",
      "Count": 336,
      "Diff": 101,
      "Percent": 30.059523
    },
    {
      "File": "staging/src/k8s.io/kubectl/pkg/util/i18n/translations/kubectl/ja_JP/LC_MESSAGES/k8s.po",
      "Count": 336,
      "Diff": 12,
      "Percent": 3.5714288
    },
    {
      "File": "staging/src/k8s.io/kubectl/pkg/util/i18n/translations/kubectl/ko_KR/LC_MESSAGES/k8s.po",
      "Count": 336,
      "Diff": 7,
      "Percent": 2.0833335
    },
    {
      "File": "staging/src/k8s.io/kubectl/pkg/util/i18n/translations/kubectl/pt_BR/LC_MESSAGES/k8s.po",
      "Count": 336,
      "Diff": 100,
      "Percent": 29.761904
    },
    {
      "File": "staging/src/k8s.io/kubectl/pkg/util/i18n/translations/kubectl/zh_CN/LC_MESSAGES/k8s.po",
      "Count": 336,
      "Diff": 101,
      "Percent": 30.059523
    },
    {
      "File": "staging/src/k8s.io/kubectl/pkg/util/i18n/translations/kubectl/zh_TW/LC_MESSAGES/k8s.po",
      "Count": 336,
      "Diff": 7,
      "Percent": 2.0833335
    },
    {
      "File": "staging/src/k8s.io/kubectl/pkg/util/i18n/translations/test/default/LC_MESSAGES/k8s.po",
      "Count": 336,
      "Diff": 2,
      "Percent": 0.5952381
    },
    {
      "File": "staging/src/k8s.io/kubectl/pkg/util/i18n/translations/test/en_US/LC_MESSAGES/k8s.po",
      "Count": 336,
      "Diff": 2,
      "Percent": 0.5952381
    }
  ]
}
```