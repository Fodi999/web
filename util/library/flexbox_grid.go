package library

func AddFlexboxGrid(cssClasses map[string]string) {
    flexboxGrid := map[string]string{
        // Flex basis
        "basis-0": "flex-basis: 0px;",
        "basis-1": "flex-basis: 0.25rem;", // 4px
        "basis-2": "flex-basis: 0.5rem;", // 8px
        "basis-3": "flex-basis: 0.75rem;", // 12px
        "basis-4": "flex-basis: 1rem;", // 16px
        "basis-5": "flex-basis: 1.25rem;", // 20px
        "basis-6": "flex-basis: 1.5rem;", // 24px
        "basis-7": "flex-basis: 1.75rem;", // 28px
        "basis-8": "flex-basis: 2rem;", // 32px
        "basis-9": "flex-basis: 2.25rem;", // 36px
        "basis-10": "flex-basis: 2.5rem;", // 40px
        "basis-11": "flex-basis: 2.75rem;", // 44px
        "basis-12": "flex-basis: 3rem;", // 48px
        "basis-14": "flex-basis: 3.5rem;", // 56px
        "basis-16": "flex-basis: 4rem;", // 64px
        "basis-20": "flex-basis: 5rem;", // 80px
        "basis-24": "flex-basis: 6rem;", // 96px
        "basis-28": "flex-basis: 7rem;", // 112px
        "basis-32": "flex-basis: 8rem;", // 128px
        "basis-36": "flex-basis: 9rem;", // 144px
        "basis-40": "flex-basis: 10rem;", // 160px
        "basis-44": "flex-basis: 11rem;", // 176px
        "basis-48": "flex-basis: 12rem;", // 192px
        "basis-52": "flex-basis: 13rem;", // 208px
        "basis-56": "flex-basis: 14rem;", // 224px
        "basis-60": "flex-basis: 15rem;", // 240px
        "basis-64": "flex-basis: 16rem;", // 256px
        "basis-72": "flex-basis: 18rem;", // 288px
        "basis-80": "flex-basis: 20rem;", // 320px
        "basis-96": "flex-basis: 24rem;", // 384px
        "basis-auto": "flex-basis: auto;",
        "basis-px": "flex-basis: 1px;",
        "basis-0.5": "flex-basis: 0.125rem;", // 2px
        "basis-1.5": "flex-basis: 0.375rem;", // 6px
        "basis-2.5": "flex-basis: 0.625rem;", // 10px
        "basis-3.5": "flex-basis: 0.875rem;", // 14px
        "basis-1/2": "flex-basis: 50%;",
        "basis-1/3": "flex-basis: 33.333333%;",
        "basis-2/3": "flex-basis: 66.666667%;",
        "basis-1/4": "flex-basis: 25%;",
        "basis-2/4": "flex-basis: 50%;",
        "basis-3/4": "flex-basis: 75%;",
        "basis-1/5": "flex-basis: 20%;",
        "basis-2/5": "flex-basis: 40%;",
        "basis-3/5": "flex-basis: 60%;",
        "basis-4/5": "flex-basis: 80%;",
        "basis-1/6": "flex-basis: 16.666667%;",
        "basis-2/6": "flex-basis: 33.333333%;",
        "basis-3/6": "flex-basis: 50%;",
        "basis-4/6": "flex-basis: 66.666667%;",
        "basis-5/6": "flex-basis: 83.333333%;",
        "basis-1/12": "flex-basis: 8.333333%;",
        "basis-2/12": "flex-basis: 16.666667%;",
        "basis-3/12": "flex-basis: 25%;",
        "basis-4/12": "flex-basis: 33.333333%;",
        "basis-5/12": "flex-basis: 41.666667%;",
        "basis-6/12": "flex-basis: 50%;",
        "basis-7/12": "flex-basis: 58.333333%;",
        "basis-8/12": "flex-basis: 66.666667%;",
        "basis-9/12": "flex-basis: 75%;",
        "basis-10/12": "flex-basis: 83.333333%;",
        "basis-11/12": "flex-basis: 91.666667%;",
        "basis-full": "flex-basis: 100%;",

        // Flex direction
        "flex-row": "flex-direction: row;",
        "flex-row-reverse": "flex-direction: row-reverse;",
        "flex-col": "flex-direction: column;",
        "flex-col-reverse": "flex-direction: column-reverse;",

        // Flex wrap
        "flex-wrap": "flex-wrap: wrap;",
        "flex-wrap-reverse": "flex-wrap: wrap-reverse;",
        "flex-nowrap": "flex-wrap: nowrap;",

        // Flex
        "flex-1": "flex: 1 1 0%;",
        "flex-auto": "flex: 1 1 auto;",
        "flex-initial": "flex: 0 1 auto;",
        "flex-none": "flex: none;",

        // Flex grow
        "grow": "flex-grow: 1;",
        "grow-0": "flex-grow: 0;",

        // Flex shrink
        "shrink": "flex-shrink: 1;",
        "shrink-0": "flex-shrink: 0;",

        // Order
        "order-1": "order: 1;",
        "order-2": "order: 2;",
        "order-3": "order: 3;",
        "order-4": "order: 4;",
        "order-5": "order: 5;",
        "order-6": "order: 6;",
        "order-7": "order: 7;",
        "order-8": "order: 8;",
        "order-9": "order: 9;",
        "order-10": "order: 10;",
        "order-11": "order: 11;",
        "order-12": "order: 12;",
        "order-first": "order: -9999;",
        "order-last": "order: 9999;",
        "order-none": "order: 0;",

        // Grid template columns
        "grid-cols-1": "grid-template-columns: repeat(1, minmax(0, 1fr));",
        "grid-cols-2": "grid-template-columns: repeat(2, minmax(0, 1fr));",
        "grid-cols-3": "grid-template-columns: repeat(3, minmax(0, 1fr));",
        "grid-cols-4": "grid-template-columns: repeat(4, minmax(0, 1fr));",
        "grid-cols-5": "grid-template-columns: repeat(5, minmax(0, 1fr));",
        "grid-cols-6": "grid-template-columns: repeat(6, minmax(0, 1fr));",
        "grid-cols-7": "grid-template-columns: repeat(7, minmax(0, 1fr));",
        "grid-cols-8": "grid-template-columns: repeat(8, minmax(0, 1fr));",
        "grid-cols-9": "grid-template-columns: repeat(9, minmax(0, 1fr));",
        "grid-cols-10": "grid-template-columns: repeat(10, minmax(0, 1fr));",
        "grid-cols-11": "grid-template-columns: repeat(11, minmax(0, 1fr));",
        "grid-cols-12": "grid-template-columns: repeat(12, minmax(0, 1fr));",
        "grid-cols-none": "grid-template-columns: none;",
        "grid-cols-subgrid": "grid-template-columns: subgrid;",

        // Grid column
        "col-auto": "grid-column: auto;",
        "col-span-1": "grid-column: span 1 / span 1;",
        "col-span-2": "grid-column: span 2 / span 2;",
        "col-span-3": "grid-column: span 3 / span 3;",
        "col-span-4": "grid-column: span 4 / span 4;",
        "col-span-5": "grid-column: span 5 / span 5;",
        "col-span-6": "grid-column: span 6 / span 6;",
        "col-span-7": "grid-column: span 7 / span 7;",
        "col-span-8": "grid-column: span 8 / span 8;",
        "col-span-9": "grid-column: span 9 / span 9;",
        "col-span-10": "grid-column: span 10 / span 10;",
        "col-span-11": "grid-column: span 11 / span 11;",
        "col-span-12": "grid-column: span 12 / span 12;",
        "col-span-full": "grid-column: 1 / -1;",
        "col-start-1": "grid-column-start: 1;",
        "col-start-2": "grid-column-start: 2;",
        "col-start-3": "grid-column-start: 3;",
        "col-start-4": "grid-column-start: 4;",
        "col-start-5": "grid-column-start: 5;",
        "col-start-6": "grid-column-start: 6;",
        "col-start-7": "grid-column-start: 7;",
        "col-start-8": "grid-column-start: 8;",
        "col-start-9": "grid-column-start: 9;",
        "col-start-10": "grid-column-start: 10;",
        "col-start-11": "grid-column-start: 11;",
        "col-start-12": "grid-column-start: 12;",
        "col-start-13": "grid-column-start: 13;",
        "col-start-auto": "grid-column-start: auto;",
        "col-end-1": "grid-column-end: 1;",
        "col-end-2": "grid-column-end: 2;",
        "col-end-3": "grid-column-end: 3;",
        "col-end-4": "grid-column-end: 4;",
        "col-end-5": "grid-column-end: 5;",
        "col-end-6": "grid-column-end: 6;",
        "col-end-7": "grid-column-end: 7;",
        "col-end-8": "grid-column-end: 8;",
        "col-end-9": "grid-column-end: 9;",
        "col-end-10": "grid-column-end: 10;",
        "col-end-11": "grid-column-end: 11;",
        "col-end-12": "grid-column-end: 12;",
        "col-end-13": "grid-column-end: 13;",
        "col-end-auto": "grid-column-end: auto;",

        // Grid template rows
        "grid-rows-1": "grid-template-rows: repeat(1, minmax(0, 1fr));",
        "grid-rows-2": "grid-template-rows: repeat(2, minmax(0, 1fr));",
        "grid-rows-3": "grid-template-rows: repeat(3, minmax(0, 1fr));",
        "grid-rows-4": "grid-template-rows: repeat(4, minmax(0, 1fr));",
        "grid-rows-5": "grid-template-rows: repeat(5, minmax(0, 1fr));",
        "grid-rows-6": "grid-template-rows: repeat(6, minmax(0, 1fr));",
        "grid-rows-7": "grid-template-rows: repeat(7, minmax(0, 1fr));",
        "grid-rows-8": "grid-template-rows: repeat(8, minmax(0, 1fr));",
        "grid-rows-9": "grid-template-rows: repeat(9, minmax(0, 1fr));",
        "grid-rows-10": "grid-template-rows: repeat(10, minmax(0, 1fr));",
        "grid-rows-11": "grid-template-rows: repeat(11, minmax(0, 1fr));",
        "grid-rows-12": "grid-template-rows: repeat(12, minmax(0, 1fr));",
        "grid-rows-none": "grid-template-rows: none;",
        "grid-rows-subgrid": "grid-template-rows: subgrid;",

        // Grid row
        "row-auto": "grid-row: auto;",
        "row-span-1": "grid-row: span 1 / span 1;",
        "row-span-2": "grid-row: span 2 / span 2;",
        "row-span-3": "grid-row: span 3 / span 3;",
        "row-span-4": "grid-row: span 4 / span 4;",
        "row-span-5": "grid-row: span 5 / span 5;",
        "row-span-6": "grid-row: span 6 / span 6;",
        "row-span-7": "grid-row: span 7 / span 7;",
        "row-span-8": "grid-row: span 8 / span 8;",
        "row-span-9": "grid-row: span 9 / span 9;",
        "row-span-10": "grid-row: span 10 / span 10;",
        "row-span-11": "grid-row: span 11 / span 11;",
        "row-span-12": "grid-row: span 12 / span 12;",
        "row-span-full": "grid-row: 1 / -1;",
        "row-start-1": "grid-row-start: 1;",
        "row-start-2": "grid-row-start: 2;",
        "row-start-3": "grid-row-start: 3;",
        "row-start-4": "grid-row-start: 4;",
        "row-start-5": "grid-row-start: 5;",
        "row-start-6": "grid-row-start: 6;",
        "row-start-7": "grid-row-start: 7;",
        "row-start-8": "grid-row-start: 8;",
        "row-start-9": "grid-row-start: 9;",
        "row-start-10": "grid-row-start: 10;",
        "row-start-11": "grid-row-start: 11;",
        "row-start-12": "grid-row-start: 12;",
        "row-start-13": "grid-row-start: 13;",
        "row-start-auto": "grid-row-start: auto;",
        "row-end-1": "grid-row-end: 1;",
        "row-end-2": "grid-row-end: 2;",
        "row-end-3": "grid-row-end: 3;",
        "row-end-4": "grid-row-end: 4;",
        "row-end-5": "grid-row-end: 5;",
        "row-end-6": "grid-row-end: 6;",
        "row-end-7": "grid-row-end: 7;",
        "row-end-8": "grid-row-end: 8;",
        "row-end-9": "grid-row-end: 9;",
        "row-end-10": "grid-row-end: 10;",
        "row-end-11": "grid-row-end: 11;",
        "row-end-12": "grid-row-end: 12;",
        "row-end-13": "grid-row-end: 13;",
        "row-end-auto": "grid-row-end: auto;",

        // Grid auto flow
        "grid-flow-row": "grid-auto-flow: row;",
        "grid-flow-col": "grid-auto-flow: column;",
        "grid-flow-dense": "grid-auto-flow: dense;",
        "grid-flow-row-dense": "grid-auto-flow: row dense;",
        "grid-flow-col-dense": "grid-auto-flow: column dense;",

        // Grid auto columns
        "auto-cols-auto": "grid-auto-columns: auto;",
        "auto-cols-min": "grid-auto-columns: min-content;",
        "auto-cols-max": "grid-auto-columns: max-content;",
        "auto-cols-fr": "grid-auto-columns: minmax(0, 1fr);",

        // Grid auto rows
        "auto-rows-auto": "grid-auto-rows: auto;",
        "auto-rows-min": "grid-auto-rows: min-content;",
        "auto-rows-max": "grid-auto-rows: max-content;",
        "auto-rows-fr": "grid-auto-rows: minmax(0, 1fr);",

        // Gap
        "gap-0": "gap: 0px;",
        "gap-x-0": "column-gap: 0px;",
        "gap-y-0": "row-gap: 0px;",
        "gap-px": "gap: 1px;",
        "gap-x-px": "column-gap: 1px;",
        "gap-y-px": "row-gap: 1px;",
        "gap-0.5": "gap: 0.125rem;", // 2px
        "gap-x-0.5": "column-gap: 0.125rem;", // 2px
        "gap-y-0.5": "row-gap: 0.125rem;", // 2px
        "gap-1": "gap: 0.25rem;", // 4px
        "gap-x-1": "column-gap: 0.25rem;", // 4px
        "gap-y-1": "row-gap: 0.25rem;", // 4px
        "gap-1.5": "gap: 0.375rem;", // 6px
        "gap-x-1.5": "column-gap: 0.375rem;", // 6px
        "gap-y-1.5": "row-gap: 0.375rem;", // 6px
        "gap-2": "gap: 0.5rem;", // 8px
        "gap-x-2": "column-gap: 0.5rem;", // 8px
        "gap-y-2": "row-gap: 0.5rem;", // 8px
        "gap-2.5": "gap: 0.625rem;", // 10px
        "gap-x-2.5": "column-gap: 0.625rem;", // 10px
        "gap-y-2.5": "row-gap: 0.625rem;", // 10px
        "gap-3": "gap: 0.75rem;", // 12px
        "gap-x-3": "column-gap: 0.75rem;", // 12px
        "gap-y-3": "row-gap: 0.75rem;", // 12px
        "gap-3.5": "gap: 0.875rem;", // 14px
        "gap-x-3.5": "column-gap: 0.875rem;", // 14px
        "gap-y-3.5": "row-gap: 0.875rem;", // 14px
        "gap-4": "gap: 1rem;", // 16px
        "gap-x-4": "column-gap: 1rem;", // 16px
        "gap-y-4": "row-gap: 1rem;", // 16px
        "gap-5": "gap: 1.25rem;", // 20px
        "gap-x-5": "column-gap: 1.25rem;", // 20px
        "gap-y-5": "row-gap: 1.25rem;", // 20px
        "gap-6": "gap: 1.5rem;", // 24px
        "gap-x-6": "column-gap: 1.5rem;", // 24px
        "gap-y-6": "row-gap: 1.5rem;", // 24px
        "gap-7": "gap: 1.75rem;", // 28px
        "gap-x-7": "column-gap: 1.75rem;", // 28px
        "gap-y-7": "row-gap: 1.75rem;", // 28px
        "gap-8": "gap: 2rem;", // 32px
        "gap-x-8": "column-gap: 2rem;", // 32px
        "gap-y-8": "row-gap: 2rem;", // 32px
        "gap-9": "gap: 2.25rem;", // 36px
        "gap-x-9": "column-gap: 2.25rem;", // 36px
        "gap-y-9": "row-gap: 2.25rem;", // 36px
        "gap-10": "gap: 2.5rem;", // 40px
        "gap-x-10": "column-gap: 2.5rem;", // 40px
        "gap-y-10": "row-gap: 2.5rem;", // 40px
        "gap-11": "gap: 2.75rem;", // 44px
        "gap-x-11": "column-gap: 2.75rem;", // 44px
        "gap-y-11": "row-gap: 2.75rem;", // 44px
        "gap-12": "gap: 3rem;", // 48px
        "gap-x-12": "column-gap: 3rem;", // 48px
        "gap-y-12": "row-gap: 3rem;", // 48px
        "gap-14": "gap: 3.5rem;", // 56px
        "gap-x-14": "column-gap: 3.5rem;", // 56px
        "gap-y-14": "row-gap: 3.5rem;", // 56px
        "gap-16": "gap: 4rem;", // 64px
        "gap-x-16": "column-gap: 4rem;", // 64px
        "gap-y-16": "row-gap: 4rem;", // 64px
        "gap-20": "gap: 5rem;", // 80px
        "gap-x-20": "column-gap: 5rem;", // 80px
        "gap-y-20": "row-gap: 5rem;", // 80px
        "gap-24": "gap: 6rem;", // 96px
        "gap-x-24": "column-gap: 6rem;", // 96px
        "gap-y-24": "row-gap: 6rem;", // 96px
        "gap-28": "gap: 7rem;", // 112px
        "gap-x-28": "column-gap: 7rem;", // 112px
        "gap-y-28": "row-gap: 7rem;", // 112px
        "gap-32": "gap: 8rem;", // 128px
        "gap-x-32": "column-gap: 8rem;", // 128px
        "gap-y-32": "row-gap: 8rem;", // 128px
        "gap-36": "gap: 9rem;", // 144px
        "gap-x-36": "column-gap: 9rem;", // 144px
        "gap-y-36": "row-gap: 9rem;", // 144px
        "gap-40": "gap: 10rem;", // 160px
        "gap-x-40": "column-gap: 10rem;", // 160px
        "gap-y-40": "row-gap: 10rem;", // 160px
        "gap-44": "gap: 11rem;", // 176px
        "gap-x-44": "column-gap: 11rem;", // 176px
        "gap-y-44": "row-gap: 11rem;", // 176px
        "gap-48": "gap: 12rem;", // 192px
        "gap-x-48": "column-gap: 12rem;", // 192px
        "gap-y-48": "row-gap: 12rem;", // 192px
        "gap-52": "gap: 13rem;", // 208px
        "gap-x-52": "column-gap: 13rem;", // 208px
        "gap-y-52": "row-gap: 13rem;", // 208px
        "gap-56": "gap: 14rem;", // 224px
        "gap-x-56": "column-gap: 14rem;", // 224px
        "gap-y-56": "row-gap: 14rem;", // 224px
        "gap-60": "gap: 15rem;", // 240px
        "gap-x-60": "column-gap: 15rem;", // 240px
        "gap-y-60": "row-gap: 15rem;", // 240px
        "gap-64": "gap: 16rem;", // 256px
        "gap-x-64": "column-gap: 16rem;", // 256px
        "gap-y-64": "row-gap: 16rem;", // 256px
        "gap-72": "gap: 18rem;", // 288px
        "gap-x-72": "column-gap: 18rem;", // 288px
        "gap-y-72": "row-gap: 18rem;", // 288px
        "gap-80": "gap: 20rem;", // 320px
        "gap-x-80": "column-gap: 20rem;", // 320px
        "gap-y-80": "row-gap: 20rem;", // 320px
        "gap-96": "gap: 24rem;", // 384px
        "gap-x-96": "column-gap: 24rem;", // 384px
        "gap-y-96": "row-gap: 24rem;", // 384px

        // Justify content
        "justify-normal": "justify-content: normal;",
        "justify-start": "justify-content: flex-start;",
        "justify-end": "justify-content: flex-end;",
        "justify-center": "justify-content: center;",
        "justify-between": "justify-content: space-between;",
        "justify-around": "justify-content: space-around;",
        "justify-evenly": "justify-content: space-evenly;",
        "justify-stretch": "justify-content: stretch;",

        // Justify items
        "justify-items-start": "justify-items: start;",
        "justify-items-end": "justify-items: end;",
        "justify-items-center": "justify-items: center;",
        "justify-items-stretch": "justify-items: stretch;",

        // Justify self
        "justify-self-auto": "justify-self: auto;",
        "justify-self-start": "justify-self: start;",
        "justify-self-end": "justify-self: end;",
        "justify-self-center": "justify-self: center;",
        "justify-self-stretch": "justify-self: stretch;",

        // Align content
        "content-normal": "align-content: normal;",
        "content-center": "align-content: center;",
        "content-start": "align-content: flex-start;",
        "content-end": "align-content: flex-end;",
        "content-between": "align-content: space-between;",
        "content-around": "align-content: space-around;",
        "content-evenly": "align-content: space-evenly;",
        "content-baseline": "align-content: baseline;",
        "content-stretch": "align-content: stretch;",

        // Align items
        "items-start": "align-items: flex-start;",
        "items-end": "align-items: flex-end;",
        "items-center": "align-items: center;",
        "items-baseline": "align-items: baseline;",
        "items-stretch": "align-items: stretch;",

        // Align self
        "self-auto": "align-self: auto;",
        "self-start": "align-self: flex-start;",
        "self-end": "align-self: flex-end;",
        "self-center": "align-self: center;",
        "self-stretch": "align-self: stretch;",
        "self-baseline": "align-self: baseline;",

        // Place content
        "place-content-center": "place-content: center;",
        "place-content-start": "place-content: start;",
        "place-content-end": "place-content: end;",
        "place-content-between": "place-content: space-between;",
        "place-content-around": "place-content: space-around;",
        "place-content-evenly": "place-content: space-evenly;",
        "place-content-baseline": "place-content: baseline;",
        "place-content-stretch": "place-content: stretch;",

        // Place items
        "place-items-start": "place-items: start;",
        "place-items-end": "place-items: end;",
        "place-items-center": "place-items: center;",
        "place-items-baseline": "place-items: baseline;",
        "place-items-stretch": "place-items: stretch;",

        // Place self
        "place-self-auto": "place-self: auto;",
        "place-self-start": "place-self: start;",
        "place-self-end": "place-self: end;",
        "place-self-center": "place-self: center;",
        "place-self-stretch": "place-self: stretch;",
    }
    for k, v := range flexboxGrid {
        cssClasses[k] = v
    }
}