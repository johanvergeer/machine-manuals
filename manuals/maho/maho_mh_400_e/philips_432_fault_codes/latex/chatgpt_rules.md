# ChatGPT Processing Rules for This Manual

This file defines specific **translation and formatting rules** for ChatGPT when processing this manual.

The pages of this manual only contains section headers and tables.

## **General Translation Rules**
- Translate all text **to English**.
- Keep technical terms consistent with machinist terminology.
- Do **not** add or invent text — only translate what is present.

## **Page-Specific Handling**
- Section headers must never be numbered, so always use `\section*{title}` instead of `\section{title}`
- Tables must always be a `tabular` wrapped inside a `table`. for example:

    ```latex
    \begin{table}[!h]
        \begin{tabular}{ll}
        O140 & Decimal point not allowed \\
        O141 & Address not allowed \\
        \end{tabular}
    \end{table}
    ```
- Always use `\begin{table}[!h]` instead of `\begin{table}[h]`
- Maintain consistent column alignment so `{ll}` for tables with two columns and `{lll}` for tables with three columns
