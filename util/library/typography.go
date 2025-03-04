package library

func AddTypography(cssClasses map[string]string) {
    typography := map[string]string{
        "font-sans": `font-family: ui-sans-serif, system-ui, sans-serif, "Apple Color Emoji", "Segoe UI Emoji", "Segoe UI Symbol", "Noto Color Emoji";`,
        "font-serif": `font-family: ui-serif, Georgia, Cambria, "Times New Roman", Times, serif;`,
        "font-mono": `font-family: ui-monospace, SFMono-Regular, Menlo, Monaco, Consolas, "Liberation Mono", "Courier New", monospace;`,
        
        "text-xs": `font-size: 0.75rem; line-height: 1rem;`,
        "text-sm": `font-size: 0.875rem; line-height: 1.25rem;`,
        "text-base": `font-size: 1rem; line-height: 1.5rem;`,
        "text-lg": `font-size: 1.125rem; line-height: 1.75rem;`,
        "text-xl": `font-size: 1.25rem; line-height: 1.75rem;`,
        "text-2xl": `font-size: 1.5rem; line-height: 2rem;`,
        "text-3xl": `font-size: 1.875rem; line-height: 2.25rem;`,
        "text-4xl": `font-size: 2.25rem; line-height: 2.5rem;`,
        "text-5xl": `font-size: 3rem; line-height: 1;`,
        "text-6xl": `font-size: 3.75rem; line-height: 1;`,
        "text-7xl": `font-size: 4.5rem; line-height: 1;`,
        "text-8xl": `font-size: 6rem; line-height: 1;`,
        "text-9xl": `font-size: 8rem; line-height: 1;`,
        
        "antialiased": `-webkit-font-smoothing: antialiased; -moz-osx-font-smoothing: grayscale;`,
        "subpixel-antialiased": `-webkit-font-smoothing: auto; -moz-osx-font-smoothing: auto;`,
        
        "italic": `font-style: italic;`,
        "not-italic": `font-style: normal;`,
        
        "font-thin": `font-weight: 100;`,
        "font-extralight": `font-weight: 200;`,
        "font-light": `font-weight: 300;`,
        "font-normal": `font-weight: 400;`,
        "font-medium": `font-weight: 500;`,
        "font-semibold": `font-weight: 600;`,
        "font-bold": `font-weight: 700;`,
        "font-extrabold": `font-weight: 800;`,
        "font-black": `font-weight: 900;`,
        
        "normal-nums": `font-variant-numeric: normal;`,
        "ordinal": `font-variant-numeric: ordinal;`,
        "slashed-zero": `font-variant-numeric: slashed-zero;`,
        "lining-nums": `font-variant-numeric: lining-nums;`,
        "oldstyle-nums": `font-variant-numeric: oldstyle-nums;`,
        "proportional-nums": `font-variant-numeric: proportional-nums;`,
        "tabular-nums": `font-variant-numeric: tabular-nums;`,
        "diagonal-fractions": `font-variant-numeric: diagonal-fractions;`,
        "stacked-fractions": `font-variant-numeric: stacked-fractions;`,
        
        "tracking-tighter": `letter-spacing: -0.05em;`,
        "tracking-tight": `letter-spacing: -0.025em;`,
        "tracking-normal": `letter-spacing: 0em;`,
        "tracking-wide": `letter-spacing: 0.025em;`,
        "tracking-wider": `letter-spacing: 0.05em;`,
        "tracking-widest": `letter-spacing: 0.1em;`,
        
      
        "line-clamp-none": `overflow: visible; display: block; -webkit-box-orient: horizontal; -webkit-line-clamp: none;`,
        
        "leading-3": `line-height: .75rem;`, /* 12px */
        "leading-4": `line-height: 1rem;`, /* 16px */
        "leading-5": `line-height: 1.25rem;`, /* 20px */
        "leading-6": `line-height: 1.5rem;`, /* 24px */
        "leading-7": `line-height: 1.75rem;`, /* 28px */
        "leading-8": `line-height: 2rem;`, /* 32px */
        "leading-9": `line-height: 2.25rem;`, /* 36px */
        "leading-10": `line-height: 2.5rem;`, /* 40px */
        "leading-none": `line-height: 1;`,
        "leading-tight": `line-height: 1.25;`,
        "leading-snug": `line-height: 1.375;`,
        "leading-normal": `line-height: 1.5;`,
        "leading-relaxed": `line-height: 1.625;`,
        "leading-loose": `line-height: 2;`,
        
        "list-image-none": `list-style-image: none;`,
        "list-inside": `list-style-position: inside;`,
        "list-outside": `list-style-position: outside;`,
        "list-none": `list-style-type: none;`,
        "list-disc": `list-style-type: disc;`,
        "list-decimal": `list-style-type: decimal;`,
        
        "text-left": `text-align: left;`,
        "text-center": `text-align: center;`,
        "text-right": `text-align: right;`,
        "text-justify": `text-align: justify;`,
        "text-start": `text-align: start;`,
        "text-end": `text-align: end;`,
        
        "underline": `text-decoration-line: underline;`,
        "overline": `text-decoration-line: overline;`,
        "line-through": `text-decoration-line: line-through;`,
        "no-underline": `text-decoration-line: none;`,
        
        "decoration-solid": `text-decoration-style: solid;`,
        "decoration-double": `text-decoration-style: double;`,
        "decoration-dotted": `text-decoration-style: dotted;`,
        "decoration-dashed": `text-decoration-style: dashed;`,
        "decoration-wavy": `text-decoration-style: wavy;`,
        
        "decoration-auto": `text-decoration-thickness: auto;`,
        "decoration-from-font": `text-decoration-thickness: from-font;`,
        "decoration-0": `text-decoration-thickness: 0px;`,
        "decoration-1": `text-decoration-thickness: 1px;`,
        "decoration-2": `text-decoration-thickness: 2px;`,
        "decoration-4": `text-decoration-thickness: 4px;`,
        "decoration-8": `text-decoration-thickness: 8px;`,
        
        "underline-offset-auto": `text-underline-offset: auto;`,
        "underline-offset-0": `text-underline-offset: 0px;`,
        "underline-offset-1": `text-underline-offset: 1px;`,
        "underline-offset-2": `text-underline-offset: 2px;`,
        "underline-offset-4": `text-underline-offset: 4px;`,
        "underline-offset-8": `text-underline-offset: 8px;`,
        
        "uppercase": `text-transform: uppercase;`,
        "lowercase": `text-transform: lowercase;`,
        "capitalize": `text-transform: capitalize;`,
        "normal-case": `text-transform: none;`,
        
     
        "text-ellipsis": `text-overflow: ellipsis;`,
        "text-clip": `text-overflow: clip;`,
        
        "text-wrap": `text-wrap: wrap;`,
        "text-nowrap": `text-wrap: nowrap;`,
        "text-balance": `text-wrap: balance;`,
        "text-pretty": `text-wrap: pretty;`,
        
        "indent-0": `text-indent: 0px;`,
        "indent-px": `text-indent: 1px;`,
        "indent-0.5": `text-indent: 0.125rem;`, /* 2px */
        "indent-1": `text-indent: 0.25rem;`, /* 4px */
        "indent-1.5": `text-indent: 0.375rem;`, /* 6px */
        "indent-2": `text-indent: 0.5rem;`, /* 8px */
        "indent-2.5": `text-indent: 0.625rem;`, /* 10px */
        "indent-3": `text-indent: 0.75rem;`, /* 12px */
        "indent-3.5": `text-indent: 0.875rem;`, /* 14px */
        "indent-4": `text-indent: 1rem;`, /* 16px */
        "indent-5": `text-indent: 1.25rem;`, /* 20px */
        "indent-6": `text-indent: 1.5rem;`, /* 24px */
        "indent-7": `text-indent: 1.75rem;`, /* 28px */
        "indent-8": `text-indent: 2rem;`, /* 32px */
        "indent-9": `text-indent: 2.25rem;`, /* 36px */
        "indent-10": `text-indent: 2.5rem;`, /* 40px */
        "indent-11": `text-indent: 2.75rem;`, /* 44px */
        "indent-12": `text-indent: 3rem;`, /* 48px */
        "indent-14": `text-indent: 3.5rem;`, /* 56px */
        "indent-16": `text-indent: 4rem;`, /* 64px */
        "indent-20": `text-indent: 5rem;`, /* 80px */
        "indent-24": `text-indent: 6rem;`, /* 96px */
        "indent-28": `text-indent: 7rem;`, /* 112px */
        "indent-32": `text-indent: 8rem;`, /* 128px */
        "indent-36": `text-indent: 9rem;`, /* 144px */
        "indent-40": `text-indent: 10rem;`, /* 160px */
        "indent-44": `text-indent: 11rem;`, /* 176px */
        "indent-48": `text-indent: 12rem;`, /* 192px */
        "indent-52": `text-indent: 13rem;`, /* 208px */
        "indent-56": `text-indent: 14rem;`, /* 224px */
        "indent-60": `text-indent: 15rem;`, /* 240px */
        "indent-64": `text-indent: 16rem;`, /* 256px */
        "indent-72": `text-indent: 18rem;`, /* 288px */
        "indent-80": `text-indent: 20rem;`, /* 320px */
        "indent-96": `text-indent: 24rem;`, /* 384px */
        
        "align-baseline": `vertical-align: baseline;`,
        "align-top": `vertical-align: top;`,
        "align-middle": `vertical-align: middle;`,
        "align-bottom": `vertical-align: bottom;`,
        "align-text-top": `vertical-align: text-top;`,
        "align-text-bottom": `vertical-align: text-bottom;`,
        "align-sub": `vertical-align: sub;`,
        "align-super": `vertical-align: super;`,
        
        "whitespace-normal": `white-space: normal;`,
        "whitespace-nowrap": `white-space: nowrap;`,
        "whitespace-pre": `white-space: pre;`,
        "whitespace-pre-line": `white-space: pre-line;`,
        "whitespace-pre-wrap": `white-space: pre-wrap;`,
        "whitespace-break-spaces": `white-space: break-spaces;`,
        
        "break-normal": `overflow-wrap: normal; word-break: normal;`,
        "break-words": `overflow-wrap: break-word;`,
        "break-all": `word-break: break-all;`,
        "break-keep": `word-break: keep-all;`,
        
        "hyphens-none": `hyphens: none;`,
        "hyphens-manual": `hyphens: manual;`,
        "hyphens-auto": `hyphens: auto;`,
        
        "content-none": `content: none;`,
    }
    for k, v := range typography {
        cssClasses[k] = v
    }
}