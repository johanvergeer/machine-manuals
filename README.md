# Machine Manual Translations

This project aims to make **older machine manuals** accessible to a broader audience by translating them into **English**. Many of these manuals are difficult to find or only available in their original language, making it challenging for non-native speakers to use them effectively.

## 🚀 Project Overview

- **Goal:** Provide clear, structured translations of machine manuals.
- **Languages:** Translated to English from various original languages.
- **Scope:** Focuses on older machines that may not have modern documentation available.

⚠ **Disclaimer:**  
I am **not a native English speaker** and **not a professional machinist**. While I have done my best to ensure accuracy, mistakes may occur. If you notice any errors or unclear translations, contributions and feedback are welcome!

---

## 📂 Project Structure

```
machine-manuals/
│── manuals/                   # Contains translated manuals
│   ├── <manufacturer_name>/    # Manufacturer folder (e.g., DMG-Mori, Schaublin)
│   │   ├── <machine_name>/     # Specific machine model (e.g., Schaublin 102, Maho MH600)
│   │   │   ├── <manual_name>/  # Name of the manual (e.g., “Service Manual”)
│   │   │   │   ├── input/      # Raw PDF scans of the original manual
│   │   │   │   ├── latex/      # Translated LaTeX files
│   │   │   │   │   ├── main.tex   # Main LaTeX file for the manual
│   │   │   │   │   ├── chapters/  # Chapters stored separately
│   │   │   │   │   ├── images/    # Images used in the manual
│
│── manual_utils/               # Utility scripts for managing manuals
│   ├── main.go                 # Main Go script
│   ├── (Other helper scripts)  # Additional utilities (TBA)
│
│── README.md                   # This file
```

---

## 📖 How to Use

### View a Manual

   - Navigate to `manuals/<manufacturer>/<machine>/<manual_name>/latex/`
   - Open `main.tex` with a LaTeX editor (e.g., Overleaf, TeXworks) and compile it.

### 📉 Optimizing PDF Size

Many forums and platforms have **file size limits**, making it difficult to share large PDFs. To reduce the size of a generated manual, you can use **Ghostscript**:

```sh
gs -sDEVICE=pdfwrite -dCompatibilityLevel=1.4 -dPDFSETTINGS=/ebook -dNOPAUSE -dQUIET -dBATCH -sOutputFile=main_compressed.pdf main.pdf
```

### What This Does
- **Compatibility:** Converts the PDF to **PDF 1.4** for wider support.
- **Size Reduction:** Uses `-dPDFSETTINGS=/ebook` to **lower resolution** while keeping text readable.
- **Automation:** Runs in **batch mode** (`-dBATCH -dNOPAUSE -dQUIET`) to suppress unnecessary output.

Example Usage

Run this command in the latex/ folder after compiling the manual:

```bash
cd manuals/Schaublin/102/operator_manual/latex/outputs
gs -sDEVICE=pdfwrite -dCompatibilityLevel=1.4 -dPDFSETTINGS=/ebook -dNOPAUSE -dQUIET -dBATCH -sOutputFile=main_compressed.pdf main.pdf
```

This will create main_compressed.pdf, a smaller version of the original PDF.

### Modify a Manual

   - Edit the `.tex` files inside the `latex/` folder.
   - If adding images, place them in the corresponding `images/` directory.

### Creating a New Manual Structure Automatically

Use the Go-based utility to quickly create a structured manual directory instead of setting it up manually.

#### **Usage**
Run the following command to generate a new manual structure:
```sh
go run main.go init_manual -m "Schaublin" -n "102" --manual_name "operator_manual"
```

**Command options:**

| Flag                | Description                                          | Default             |
|---------------------|--------------------------------------------------|---------------------|
| `-m, --manufacturer` | The machine manufacturer (e.g., "Schaublin", "Maho") | **(Required)** |
| `-n, --machine_name` | The machine model (e.g., "102", "MH600")            | **(Required)** |
| `--manual_name`      | The type of manual (e.g., "service_manual", "operator_manual") | `"operator_manual"` |

**This will create the following structure:**

```
manuals/Schaublin/102/operator_manual/
│── input/      # Raw PDFs
│── latex/      # Translated LaTeX files
│   ├── main.tex
│   ├── chapters/
│   ├── images/
```

### 📜 ChatGPT Rules for Manual Processing

Each manual includes a **`chatgpt_rules.md`** file inside its `latex/` folder. This file contains specific **translation and formatting rules** to ensure consistency when using ChatGPT for document processing.

#### **Purpose of `chatgpt_rules.md`**
- Defines **translation guidelines** (e.g., "Translate everything to English").
- Specifies **formatting conventions** (e.g., "Convert underlined text to `\textbf{}`").
- Provides **page-specific handling instructions** (e.g., "Convert tables to `tabularx` in LaTeX").
- Can be **customized per manual** to handle special cases.

#### **Default Rules**
When a new manual is initialized using `init_manual`, the following default rules are added:

```markdown
# ChatGPT Processing Rules for This Manual

This file defines specific **translation and formatting rules** for ChatGPT when processing this manual.

## **General Translation Rules**
- Translate all text **to English**.
- Keep technical terms consistent with machinist terminology.
- Do **not** add or invent text — only translate what is present.

## **Formatting Conventions**
- Convert **underlined text** to `\textbf{}` (bold) in LaTeX.

## **Page-Specific Handling**
- If a page contains only a table, convert it to a `tabularx` LaTeX table.
- If a page has hand-drawn annotations, **omit them** unless they contain relevant data.
```

This file can be modified for **manual-specific instructions**.

#### **Where to Find It**
- Located inside each manual’s **`latex/`** folder.
- Can be **manually edited** to add specific instructions for complex manuals.

This ensures that all manuals follow **consistent processing rules** when using ChatGPT.

## 📷 Converting PDFs to Images

Some machine manuals contain **scanned pages** or **diagrams** that need to be extracted as images for better readability or inclusion in the LaTeX document. To automate this, use the `pdf_to_jpg` command.

### **Usage**

Saves the extracted images in input/images/.

```
mycli pdf_to_jpg -m "Schaublin" -n "102" --manual_name "operator_manual"
```

### **Command Options**

| Flag                   | Description                                              | Default               |
|------------------------|----------------------------------------------------------|-----------------------|
| `-m, --manufacturer`   | The machine manufacturer (e.g., "Schaublin", "Maho")    | **(Required)**        |
| `-n, --machine_name`   | The machine model (e.g., "102", "MH600")                | **(Required)**        |
| `--manual_name`        | The type of manual (e.g., "service_manual", "operator_manual") | `"operator_manual"` |

### **How It Works**
- Takes all PDFs from **`input/raw_pdf/`**.
- Converts each page to a **JPG file**.
- Saves the extracted images in **`input/images/`**.

---

## 🔧 Utility Scripts (manual_utils)

In addition to init_manual, more Go-based command-line tools are planned to help maintain consistency across manuals. These tools will automate tasks like:

- Ensuring correct folder structures.
- Detecting missing or unused files.
- Formatting LaTeX content to match project guidelines.


📌 **More details about these utilities will be added later.**

---

## 🛠 Contributions & Feedback

Contributions are welcome! If you:
- Find errors in a translation,
- Have suggestions for improving terminology,
- Want to help with formatting,

... feel free to **open an issue or submit a pull request**.

---

## 📜 License

This project is **not affiliated with any machine manufacturers**. It is a **community-driven effort** to make manuals more accessible.  

⚠ **Note:** If a manufacturer requests the removal of a manual for legal reasons, it will be taken down.

---

## ⭐ Acknowledgments

A big thanks to everyone who has contributed knowledge, corrections, or feedback. This project exists to help machinists, restorers, and enthusiasts keep these classic machines running!
