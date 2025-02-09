# ChatGPT Processing Rules for This Manual

This file defines specific **translation and formatting rules** for ChatGPT when processing this manual.

## General Translation Rules
- Translate all text **to English**.
- Keep technical terms consistent with machinist terminology.
- Do **not** add or invent text — only translate what is present.
- You may **reformat text** (e.g., using lists) if it improves readability.
- Always use **jargon an American machinist would understand**.
- The full **LaTeX preamble is not needed**, as this document already exists.

## Formatting Conventions
- Convert **underlined text** (except headers) to `\textbf{}` (bold).
- Other text can be bolded **if it adds clarity** for the reader.

## Chapter and Section Hierarchy
Chapters and sections always have a **numbered prefix**:

| Prefix Format | Meaning |
|--------------|---------|
| `1` digit    | **Chapter** |
| `2` digits   | **Section** |
| `3` digits   | **Subsection** |
| No digits    | **Not a chapter or section** |

## Notes, Remarks, and Examples
To ensure consistent formatting for notes, procedures, and examples, use the following commands:

| Purpose  | LaTeX Command |
|----------|--------------|
| **Notes (remarks)** | `\notes` |
| **Procedures** | `\procedure` |
| **Examples** | `\example` |

These commands **do not take parameters**.  

**Example Usage:**

\notes

For machines with an automatic tool changer, when executing a block with the tool change instruction \textbf{M6},  
no program interruption occurs. The tool is automatically changed after automatic retraction into the tool change position.

## Image Handling

### Icons in Left Margin
For list items with images in the left margin, use the `\iconitem{}` command inside an `itemize` environment:

\begin{itemize}
    \iconitem{Press the \textbf{MENU} button.}{menu.jpg}
\end{itemize}

- An `\iconitem{}` **can contain multiple images**, which will be positioned next to each other in the final render.
- **Spacing Rule for Consecutive `\iconitem{}` Entries:**
  - If the text inside each `\iconitem{}` is too short to fill at least **two lines**, add `\vspace{.5cm}` between them.
  - If the text is **long enough** to create spacing naturally, **no extra spacing is needed**.

**Example Usage:**

\begin{itemize}
    \iconitem{Press the \textbf{TEACH IN}, \textbf{SINGLE}, or \textbf{AUTO} button according to the signal displayed on line 1 of the screen (TEACH IN, SINGLE, or AUTO).}{teach_in.jpg, single.jpg, auto.jpg}
\end{itemize}
% No extra spacing needed
\begin{itemize}
    \iconitem{Press the \textbf{PROG.MEM} button.}{prog_mem.jpg}
\end{itemize}
\vspace{.5cm}
\begin{itemize}
    \iconitem{Press the \textbf{MENU} button.}{menu.jpg}
\end{itemize}

### Other Images
#### 1. Logical and Consistent Naming
- Use descriptive filenames that reflect the image content.
- **Avoid generic names** (e.g., `image1.jpg`), use meaningful names like `program_list.jpg` or `data_transfer.jpg`.

#### 2. Reusing Images
- If an image appears multiple times in the document, **reuse the same filename** to avoid duplicates.

#### 3. Image Storage and Path
- All images must be stored in the **`images/`** directory.
- Since this directory is **already configured in LaTeX**, do **not** include `images/` in the file path.

#### 4. Image Insertion Format
- Always center images using `\begin{center}...\end{center}`.
- Use `\includegraphics[width=0.6\linewidth]{image_name.jpg}` for consistency.

**Example Usage:**

\begin{center}
    \includegraphics[width=0.6\linewidth]{program_entry.jpg}
\end{center}
