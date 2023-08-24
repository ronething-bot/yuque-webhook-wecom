<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->
**Table of Contents**  *generated with [DocToc](https://github.com/thlorenz/doctoc)*

- [yuque-webhook-wecom](#yuque-webhook-wecom)
  - [Introduction](#introduction)
  - [Configuration](#configuration)
  - [Effect](#effect)
  - [Acknowledgements](#acknowledgements)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

## yuque-webhook-wecom

### Introduction

~~Yuque webhook by default only supports dingtalk. To support WeChat Work bots, one has to write their own code for the corresponding parsing and pushing.~~

Actually, it is supported. It's just that it seems I couldn't find it in the Yuque documentation. Once configured, it can be sent normally. However, it's in card format. If you want to customize the format and such, you can do it through this program.

### Configuration

- [sample.yaml](./config/sample.yaml)

### Effect

![Effect](./asserts/sample.png)

### Acknowledgements

- [go.mod](go.mod)