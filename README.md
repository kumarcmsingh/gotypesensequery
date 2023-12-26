# gotypesensequery

[![Go Report Card](https://goreportcard.com/badge/github.com/kumarcmsingh/gotypesensequery)](https://goreportcard.com/report/github.com/kumarcmsingh/gotypesensequery)

**gotypesensequery** is a Go package designed to simplify the generation of Typesense-supported query strings. It provides a versatile set of utilities for constructing filter conditions across different field types, making it easier for developers to build complex and precise Typesense queries in their Go applications.

## Table of Contents

- [Features](#features)
- [Usage](#usage)
  - [Installation](#installation)
  - [Example](#example)
- [Functionality](#functionality)
  - [Text Field Conditions](#text-field-conditions)
  - [Number Field Conditions](#number-field-conditions)
  - [Date and DateTime Field Conditions](#date-and-datetime-field-conditions)
- [Human-Readable Date Conversions](#human-readable-date-conversions)
- [Contributing](#contributing)
- [License](#license)

## Features

- **Effortless Query Construction:** Easily create Typesense queries by composing filter conditions.
- **Versatile Field Type Support:** Handle various field types, including text, number, date, and datetime.
- **Human-Readable Date Conversion:** Convert human-readable date inputs to Typesense datetime conditions seamlessly.

## Usage

### Installation

```bash
go get -u github.com/kumarcmsingh/gotypesensequery
