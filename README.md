feature：

- 去重
- 过滤 params

simple stdin use

```bash
cat katana_result.txt | anewp | uro | nuclei -t ~/nuclei-tripse/dast/ -dast -lfa | tee dast_template_result.txt
```

