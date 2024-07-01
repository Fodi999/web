package library

func AddSpacing(cssClasses map[string]string) {
    spacing := map[string]string{
        // Padding
        "p-0": "padding: 0px;",
        "px-0": "padding-left: 0px; padding-right: 0px;",
        "py-0": "padding-top: 0px; padding-bottom: 0px;",
       
        "pt-0": "padding-top: 0px;",
        "pr-0": "padding-right: 0px;",
        "pb-0": "padding-bottom: 0px;",
        "pl-0": "padding-left: 0px;",
        "p-px": "padding: 1px;",
        "px-px": "padding-left: 1px; padding-right: 1px;",
        "py-px": "padding-top: 1px; padding-bottom: 1px;",
       
        "pt-px": "padding-top: 1px;",
        "pr-px": "padding-right: 1px;",
        "pb-px": "padding-bottom: 1px;",
        "pl-px": "padding-left: 1px;",
        "p-0.5": "padding: 0.125rem;", // 2px
        "px-0.5": "padding-left: 0.125rem; padding-right: 0.125rem;", // 2px
        "py-0.5": "padding-top: 0.125rem; padding-bottom: 0.125rem;", // 2px
       
        "pt-0.5": "padding-top: 0.125rem;", // 2px
        "pr-0.5": "padding-right: 0.125rem;", // 2px
        "pb-0.5": "padding-bottom: 0.125rem;", // 2px
        "pl-0.5": "padding-left: 0.125rem;", // 2px
        "p-1": "padding: 0.25rem;", // 4px
        "px-1": "padding-left: 0.25rem; padding-right: 0.25rem;", // 4px
        "py-1": "padding-top: 0.25rem; padding-bottom: 0.25rem;", // 4px
        
        "pt-1": "padding-top: 0.25rem;", // 4px
        "pr-1": "padding-right: 0.25rem;", // 4px
        "pb-1": "padding-bottom: 0.25rem;", // 4px
        "pl-1": "padding-left: 0.25rem;", // 4px
        "p-1.5": "padding: 0.375rem;", // 6px
        "px-1.5": "padding-left: 0.375rem; padding-right: 0.375rem;", // 6px
        "py-1.5": "padding-top: 0.375rem; padding-bottom: 0.375rem;", // 6px
       
        "pt-1.5": "padding-top: 0.375rem;", // 6px
        "pr-1.5": "padding-right: 0.375rem;", // 6px
        "pb-1.5": "padding-bottom: 0.375rem;", // 6px
        "pl-1.5": "padding-left: 0.375rem;", // 6px
        "p-2": "padding: 0.5rem;", // 8px
        "px-2": "padding-left: 0.5rem; padding-right: 0.5rem;", // 8px
        "py-2": "padding-top: 0.5rem; padding-bottom: 0.5rem;", // 8px
        
        "pt-2": "padding-top: 0.5rem;", // 8px
        "pr-2": "padding-right: 0.5rem;", // 8px
        "pb-2": "padding-bottom: 0.5rem;", // 8px
        "pl-2": "padding-left: 0.5rem;", // 8px
        "p-2.5": "padding: 0.625rem;", // 10px
        "px-2.5": "padding-left: 0.625rem; padding-right: 0.625rem;", // 10px
        "py-2.5": "padding-top: 0.625rem; padding-bottom: 0.625rem;", // 10px
        
        "pt-2.5": "padding-top: 0.625rem;", // 10px
        "pr-2.5": "padding-right: 0.625rem;", // 10px
        "pb-2.5": "padding-bottom: 0.625rem;", // 10px
        "pl-2.5": "padding-left: 0.625rem;", // 10px
        "p-3": "padding: 0.75rem;", // 12px
        "px-3": "padding-left: 0.75rem; padding-right: 0.75rem;", // 12px
        "py-3": "padding-top: 0.75rem; padding-bottom: 0.75rem;", // 12px
       
        "pt-3": "padding-top: 0.75rem;", // 12px
        "pr-3": "padding-right: 0.75rem;", // 12px
        "pb-3": "padding-bottom: 0.75rem;", // 12px
        "pl-3": "padding-left: 0.75rem;", // 12px
        "p-3.5": "padding: 0.875rem;", // 14px
        "px-3.5": "padding-left: 0.875rem; padding-right: 0.875rem;", // 14px
        "py-3.5": "padding-top: 0.875rem; padding-bottom: 0.875rem;", // 14px
       
        "pt-3.5": "padding-top: 0.875rem;", // 14px
        "pr-3.5": "padding-right: 0.875rem;", // 14px
        "pb-3.5": "padding-bottom: 0.875rem;", // 14px
        "pl-3.5": "padding-left: 0.875rem;", // 14px
        "p-4": "padding: 1rem;", // 16px
        "px-4": "padding-left: 1rem; padding-right: 1rem;", // 16px
        "py-4": "padding-top: 1rem; padding-bottom: 1rem;", // 16px
     
        "pt-4": "padding-top: 1rem;", // 16px
        "pr-4": "padding-right: 1rem;", // 16px
        "pb-4": "padding-bottom: 1rem;", // 16px
        "pl-4": "padding-left: 1rem;", // 16px
        "p-5": "padding: 1.25rem;", // 20px
        "px-5": "padding-left: 1.25rem; padding-right: 1.25rem;", // 20px
        "py-5": "padding-top: 1.25rem; padding-bottom: 1.25rem;", // 20px
       
        "pt-5": "padding-top: 1.25rem;", // 20px
        "pr-5": "padding-right: 1.25rem;", // 20px
        "pb-5": "padding-bottom: 1.25rem;", // 20px
        "pl-5": "padding-left: 1.25rem;", // 20px
        "p-6": "padding: 1.5rem;", // 24px
        "px-6": "padding-left: 1.5rem; padding-right: 1.5rem;", // 24px
        "py-6": "padding-top: 1.5rem; padding-bottom: 1.5rem;", // 24px
       
        "pt-6": "padding-top: 1.5rem;", // 24px
        "pr-6": "padding-right: 1.5rem;", // 24px
        "pb-6": "padding-bottom: 1.5rem;", // 24px
        "pl-6": "padding-left: 1.5rem;", // 24px
        "p-7": "padding: 1.75rem;", // 28px
        "px-7": "padding-left: 1.75rem; padding-right: 1.75rem;", // 28px
        "py-7": "padding-top: 1.75rem; padding-bottom: 1.75rem;", // 28px
      
        "pt-7": "padding-top: 1.75rem;", // 28px
        "pr-7": "padding-right: 1.75rem;", // 28px
        "pb-7": "padding-bottom: 1.75rem;", // 28px
        "pl-7": "padding-left: 1.75rem;", // 28px
        "p-8": "padding: 2rem;", // 32px
        "px-8": "padding-left: 2rem; padding-right: 2rem;", // 32px
        "py-8": "padding-top: 2rem; padding-bottom: 2rem;", // 32px
      
        "pt-8": "padding-top: 2rem;", // 32px
        "pr-8": "padding-right: 2rem;", // 32px
        "pb-8": "padding-bottom: 2rem;", // 32px
        "pl-8": "padding-left: 2rem;", // 32px
        "p-9": "padding: 2.25rem;", // 36px
        "px-9": "padding-left: 2.25rem; padding-right: 2.25rem;", // 36px
        "py-9": "padding-top: 2.25rem; padding-bottom: 2.25rem;", // 36px
       
        "pt-9": "padding-top: 2.25rem;", // 36px
        "pr-9": "padding-right: 2.25rem;", // 36px
        "pb-9": "padding-bottom: 2.25rem;", // 36px
        "pl-9": "padding-left: 2.25rem;", // 36px
        "p-10": "padding: 2.5rem;", // 40px
        "px-10": "padding-left: 2.5rem; padding-right: 2.5rem;", // 40px
        "py-10": "padding-top: 2.5rem; padding-bottom: 2.5rem;", // 40px
      
        "pt-10": "padding-top: 2.5rem;", // 40px
        "pr-10": "padding-right: 2.5rem;", // 40px
        "pb-10": "padding-bottom: 2.5rem;", // 40px
        "pl-10": "padding-left: 2.5rem;", // 40px
        "p-11": "padding: 2.75rem;", // 44px
        "px-11": "padding-left: 2.75rem; padding-right: 2.75rem;", // 44px
        "py-11": "padding-top: 2.75rem; padding-bottom: 2.75rem;", // 44px
        
        "pt-11": "padding-top: 2.75rem;", // 44px
        "pr-11": "padding-right: 2.75rem;", // 44px
        "pb-11": "padding-bottom: 2.75rem;", // 44px
        "pl-11": "padding-left: 2.75rem;", // 44px
        "p-12": "padding: 3rem;", // 48px
        "px-12": "padding-left: 3rem; padding-right: 3rem;", // 48px
        "py-12": "padding-top: 3rem; padding-bottom: 3rem;", // 48px
      
        "pt-12": "padding-top: 3rem;", // 48px
        "pr-12": "padding-right: 3rem;", // 48px
        "pb-12": "padding-bottom: 3rem;", // 48px
        "pl-12": "padding-left: 3rem;", // 48px
        "p-14": "padding: 3.5rem;", // 56px
        "px-14": "padding-left: 3.5rem; padding-right: 3.5rem;", // 56px
        "py-14": "padding-top: 3.5rem; padding-bottom: 3.5rem;", // 56px
      
        "pt-14": "padding-top: 3.5rem;", // 56px
        "pr-14": "padding-right: 3.5rem;", // 56px
        "pb-14": "padding-bottom: 3.5rem;", // 56px
        "pl-14": "padding-left: 3.5rem;", // 56px
        "p-16": "padding: 4rem;", // 64px
        "px-16": "padding-left: 4rem; padding-right: 4rem;", // 64px
        "py-16": "padding-top: 4rem; padding-bottom: 4rem;", // 64px
       
        "pt-16": "padding-top: 4rem;", // 64px
        "pr-16": "padding-right: 4rem;", // 64px
        "pb-16": "padding-bottom: 4rem;", // 64px
        "pl-16": "padding-left: 4rem;", // 64px
        "p-20": "padding: 5rem;", // 80px
        "px-20": "padding-left: 5rem; padding-right: 5rem;", // 80px
        "py-20": "padding-top: 5rem; padding-bottom: 5rem;", // 80px
      
        "pt-20": "padding-top: 5rem;", // 80px
        "pr-20": "padding-right: 5rem;", // 80px
        "pb-20": "padding-bottom: 5rem;", // 80px
        "pl-20": "padding-left: 5rem;", // 80px
        "p-24": "padding: 6rem;", // 96px
        "px-24": "padding-left: 6rem; padding-right: 6rem;", // 96px
        "py-24": "padding-top: 6rem; padding-bottom: 6rem;", // 96px
     
        "pt-24": "padding-top: 6rem;", // 96px
        "pr-24": "padding-right: 6rem;", // 96px
        "pb-24": "padding-bottom: 6rem;", // 96px
        "pl-24": "padding-left: 6rem;", // 96px
        "p-28": "padding: 7rem;", // 112px
        "px-28": "padding-left: 7rem; padding-right: 7rem;", // 112px
        "py-28": "padding-top: 7rem; padding-bottom: 7rem;", // 112px
       
        "pt-28": "padding-top: 7rem;", // 112px
        "pr-28": "padding-right: 7rem;", // 112px
        "pb-28": "padding-bottom: 7rem;", // 112px
        "pl-28": "padding-left: 7rem;", // 112px
        "p-32": "padding: 8rem;", // 128px
        "px-32": "padding-left: 8rem; padding-right: 8rem;", // 128px
        "py-32": "padding-top: 8rem; padding-bottom: 8rem;", // 128px
       
        "pt-32": "padding-top: 8rem;", // 128px
        "pr-32": "padding-right: 8rem;", // 128px
        "pb-32": "padding-bottom: 8rem;", // 128px
        "pl-32": "padding-left: 8rem;", // 128px
        "p-36": "padding: 9rem;", // 144px
        "px-36": "padding-left: 9rem; padding-right: 9rem;", // 144px
        "py-36": "padding-top: 9rem; padding-bottom: 9rem;", // 144px
       
        "pt-36": "padding-top: 9rem;", // 144px
        "pr-36": "padding-right: 9rem;", // 144px
        "pb-36": "padding-bottom: 9rem;", // 144px
        "pl-36": "padding-left: 9rem;", // 144px
        "p-40": "padding: 10rem;", // 160px
        "px-40": "padding-left: 10rem; padding-right: 10rem;", // 160px
        "py-40": "padding-top: 10rem; padding-bottom: 10rem;", // 160px
       
        "pt-40": "padding-top: 10rem;", // 160px
        "pr-40": "padding-right: 10rem;", // 160px
        "pb-40": "padding-bottom: 10rem;", // 160px
        "pl-40": "padding-left: 10rem;", // 160px
        "p-44": "padding: 11rem;", // 176px
        "px-44": "padding-left: 11rem; padding-right: 11rem;", // 176px
        "py-44": "padding-top: 11rem; padding-bottom: 11rem;", // 176px
      
        "pt-44": "padding-top: 11rem;", // 176px
        "pr-44": "padding-right: 11rem;", // 176px
        "pb-44": "padding-bottom: 11rem;", // 176px
        "pl-44": "padding-left: 11rem;", // 176px
        "p-48": "padding: 12rem;", // 192px
        "px-48": "padding-left: 12rem; padding-right: 12rem;", // 192px
        "py-48": "padding-top: 12rem; padding-bottom: 12rem;", // 192px
       
        "pt-48": "padding-top: 12rem;", // 192px
        "pr-48": "padding-right: 12rem;", // 192px
        "pb-48": "padding-bottom: 12rem;", // 192px
        "pl-48": "padding-left: 12rem;", // 192px
        "p-52": "padding: 13rem;", // 208px
        "px-52": "padding-left: 13rem; padding-right: 13rem;", // 208px
        "py-52": "padding-top: 13rem; padding-bottom: 13rem;", // 208px
       
        "pt-52": "padding-top: 13rem;", // 208px
        "pr-52": "padding-right: 13rem;", // 208px
        "pb-52": "padding-bottom: 13rem;", // 208px
        "pl-52": "padding-left: 13rem;", // 208px
        "p-56": "padding: 14rem;", // 224px
        "px-56": "padding-left: 14rem; padding-right: 14rem;", // 224px
        "py-56": "padding-top: 14rem; padding-bottom: 14rem;", // 224px
      
        "pt-56": "padding-top: 14rem;", // 224px
        "pr-56": "padding-right: 14rem;", // 224px
        "pb-56": "padding-bottom: 14rem;", // 224px
        "pl-56": "padding-left: 14rem;", // 224px
        "p-60": "padding: 15rem;", // 240px
        "px-60": "padding-left: 15rem; padding-right: 15rem;", // 240px
        "py-60": "padding-top: 15rem; padding-bottom: 15rem;", // 240px
       
        "pt-60": "padding-top: 15rem;", // 240px
        "pr-60": "padding-right: 15rem;", // 240px
        "pb-60": "padding-bottom: 15rem;", // 240px
        "pl-60": "padding-left: 15rem;", // 240px
        "p-64": "padding: 16rem;", // 256px
        "px-64": "padding-left: 16rem; padding-right: 16rem;", // 256px
        "py-64": "padding-top: 16rem; padding-bottom: 16rem;", // 256px
      
        "pt-64": "padding-top: 16rem;", // 256px
        "pr-64": "padding-right: 16rem;", // 256px
        "pb-64": "padding-bottom: 16rem;", // 256px
        "pl-64": "padding-left: 16rem;", // 256px
        "p-72": "padding: 18rem;", // 288px
        "px-72": "padding-left: 18rem; padding-right: 18rem;", // 288px
        "py-72": "padding-top: 18rem; padding-bottom: 18rem;", // 288px
      
        "pt-72": "padding-top: 18rem;", // 288px
        "pr-72": "padding-right: 18rem;", // 288px
        "pb-72": "padding-bottom: 18rem;", // 288px
        "pl-72": "padding-left: 18rem;", // 288px
        "p-80": "padding: 20rem;", // 320px
        "px-80": "padding-left: 20rem; padding-right: 20rem;", // 320px
        "py-80": "padding-top: 20rem; padding-bottom: 20rem;", // 320px
       
        "pt-80": "padding-top: 20rem;", // 320px
        "pr-80": "padding-right: 20rem;", // 320px
        "pb-80": "padding-bottom: 20rem;", // 320px
        "pl-80": "padding-left: 20rem;", // 320px
        "p-96": "padding: 24rem;", // 384px
        "px-96": "padding-left: 24rem; padding-right: 24rem;", // 384px
        "py-96": "padding-top: 24rem; padding-bottom: 24rem;", // 384px
       
        "pt-96": "padding-top: 24rem;", // 384px
        "pr-96": "padding-right: 24rem;", // 384px
        "pb-96": "padding-bottom: 24rem;", // 384px
        "pl-96": "padding-left: 24rem;", // 384px

        // Margin
        "m-0": "margin: 0px;",
        "mx-0": "margin-left: 0px; margin-right: 0px;",
        "my-0": "margin-top: 0px; margin-bottom: 0px;",
       
        "mt-0": "margin-top: 0px;",
        "mr-0": "margin-right: 0px;",
        "mb-0": "margin-bottom: 0px;",
        "ml-0": "margin-left: 0px;",
        "m-px": "margin: 1px;",
        "mx-px": "margin-left: 1px; margin-right: 1px;",
        "my-px": "margin-top: 1px; margin-bottom: 1px;",
      
        "mt-px": "margin-top: 1px;",
        "mr-px": "margin-right: 1px;",
        "mb-px": "margin-bottom: 1px;",
        "ml-px": "margin-left: 1px;",
        "m-0.5": "margin: 0.125rem;", // 2px
        "mx-0.5": "margin-left: 0.125rem; margin-right: 0.125rem;", // 2px
        "my-0.5": "margin-top: 0.125rem; margin-bottom: 0.125rem;", // 2px
        
        "mt-0.5": "margin-top: 0.125rem;", // 2px
        "mr-0.5": "margin-right: 0.125rem;", // 2px
        "mb-0.5": "margin-bottom: 0.125rem;", // 2px
        "ml-0.5": "margin-left: 0.125rem;", // 2px
        "m-1": "margin: 0.25rem;", // 4px
        "mx-1": "margin-left: 0.25rem; margin-right: 0.25rem;", // 4px
        "my-1": "margin-top: 0.25rem; margin-bottom: 0.25rem;", // 4px
       
        "mt-1": "margin-top: 0.25rem;", // 4px
        "mr-1": "margin-right: 0.25rem;", // 4px
        "mb-1": "margin-bottom: 0.25rem;", // 4px
        "ml-1": "margin-left: 0.25rem;", // 4px
        "m-1.5": "margin: 0.375rem;", // 6px
        "mx-1.5": "margin-left: 0.375rem; margin-right: 0.375rem;", // 6px
        "my-1.5": "margin-top: 0.375rem; margin-bottom: 0.375rem;", // 6px
      
        "mt-1.5": "margin-top: 0.375rem;", // 6px
        "mr-1.5": "margin-right: 0.375rem;", // 6px
        "mb-1.5": "margin-bottom: 0.375rem;", // 6px
        "ml-1.5": "margin-left: 0.375rem;", // 6px
        "m-2": "margin: 0.5rem;", // 8px
        "mx-2": "margin-left: 0.5rem; margin-right: 0.5rem;", // 8px
        "my-2": "margin-top: 0.5rem; margin-bottom: 0.5rem;", // 8px
       
        "mt-2": "margin-top: 0.5rem;", // 8px
        "mr-2": "margin-right: 0.5rem;", // 8px
        "mb-2": "margin-bottom: 0.5rem;", // 8px
        "ml-2": "margin-left: 0.5rem;", // 8px
        "m-2.5": "margin: 0.625rem;", // 10px
        "mx-2.5": "margin-left: 0.625rem; margin-right: 0.625rem;", // 10px
        "my-2.5": "margin-top: 0.625rem; margin-bottom: 0.625rem;", // 10px
      
        "mt-2.5": "margin-top: 0.625rem;", // 10px
        "mr-2.5": "margin-right: 0.625rem;", // 10px
        "mb-2.5": "margin-bottom: 0.625rem;", // 10px
        "ml-2.5": "margin-left: 0.625rem;", // 10px
        "m-3": "margin: 0.75rem;", // 12px
        "mx-3": "margin-left: 0.75rem; margin-right: 0.75rem;", // 12px
        "my-3": "margin-top: 0.75rem; margin-bottom: 0.75rem;", // 12px
       
        "mt-3": "margin-top: 0.75rem;", // 12px
        "mr-3": "margin-right: 0.75rem;", // 12px
        "mb-3": "margin-bottom: 0.75rem;", // 12px
        "ml-3": "margin-left: 0.75rem;", // 12px
        "m-3.5": "margin: 0.875rem;", // 14px
        "mx-3.5": "margin-left: 0.875rem; margin-right: 0.875rem;", // 14px
        "my-3.5": "margin-top: 0.875rem; margin-bottom: 0.875rem;", // 14px
     
        "mt-3.5": "margin-top: 0.875rem;", // 14px
        "mr-3.5": "margin-right: 0.875rem;", // 14px
        "mb-3.5": "margin-bottom: 0.875rem;", // 14px
        "ml-3.5": "margin-left: 0.875rem;", // 14px
        "m-4": "margin: 1rem;", // 16px
        "mx-4": "margin-left: 1rem; margin-right: 1rem;", // 16px
        "my-4": "margin-top: 1rem; margin-bottom: 1rem;", // 16px
   
        "mt-4": "margin-top: 1rem;", // 16px
        "mr-4": "margin-right: 1rem;", // 16px
        "mb-4": "margin-bottom: 1rem;", // 16px
        "ml-4": "margin-left: 1rem;", // 16px
        "m-5": "margin: 1.25rem;", // 20px
        "mx-5": "margin-left: 1.25rem; margin-right: 1.25rem;", // 20px
        "my-5": "margin-top: 1.25rem; margin-bottom: 1.25rem;", // 20px
       
        "mt-5": "margin-top: 1.25rem;", // 20px
        "mr-5": "margin-right: 1.25rem;", // 20px
        "mb-5": "margin-bottom: 1.25rem;", // 20px
        "ml-5": "margin-left: 1.25rem;", // 20px
        "m-6": "margin: 1.5rem;", // 24px
        "mx-6": "margin-left: 1.5rem; margin-right: 1.5rem;", // 24px
        "my-6": "margin-top: 1.5rem; margin-bottom: 1.5rem;", // 24px
       
        "mt-6": "margin-top: 1.5rem;", // 24px
        "mr-6": "margin-right: 1.5rem;", // 24px
        "mb-6": "margin-bottom: 1.5rem;", // 24px
        "ml-6": "margin-left: 1.5rem;", // 24px
        "m-7": "margin: 1.75rem;", // 28px
        "mx-7": "margin-left: 1.75rem; margin-right: 1.75rem;", // 28px
        "my-7": "margin-top: 1.75rem; margin-bottom: 1.75rem;", // 28px
     
        "mt-7": "margin-top: 1.75rem;", // 28px
        "mr-7": "margin-right: 1.75rem;", // 28px
        "mb-7": "margin-bottom: 1.75rem;", // 28px
        "ml-7": "margin-left: 1.75rem;", // 28px
        "m-8": "margin: 2rem;", // 32px
        "mx-8": "margin-left: 2rem; margin-right: 2rem;", // 32px
        "my-8": "margin-top: 2rem; margin-bottom: 2rem;", // 32px
       
        "mt-8": "margin-top: 2rem;", // 32px
        "mr-8": "margin-right: 2rem;", // 32px
        "mb-8": "margin-bottom: 2rem;", // 32px
        "ml-8": "margin-left: 2rem;", // 32px
        "m-9": "margin: 2.25rem;", // 36px
        "mx-9": "margin-left: 2.25rem; margin-right: 2.25rem;", // 36px
        "my-9": "margin-top: 2.25rem; margin-bottom: 2.25rem;", // 36px
       
        "mt-9": "margin-top: 2.25rem;", // 36px
        "mr-9": "margin-right: 2.25rem;", // 36px
        "mb-9": "margin-bottom: 2.25rem;", // 36px
        "ml-9": "margin-left: 2.25rem;", // 36px
        "m-10": "margin: 2.5rem;", // 40px
        "mx-10": "margin-left: 2.5rem; margin-right: 2.5rem;", // 40px
        "my-10": "margin-top: 2.5rem; margin-bottom: 2.5rem;", // 40px
     
        "mt-10": "margin-top: 2.5rem;", // 40px
        "mr-10": "margin-right: 2.5rem;", // 40px
        "mb-10": "margin-bottom: 2.5rem;", // 40px
        "ml-10": "margin-left: 2.5rem;", // 40px
        "m-11": "margin: 2.75rem;", // 44px
        "mx-11": "margin-left: 2.75rem; margin-right: 2.75rem;", // 44px
        "my-11": "margin-top: 2.75rem; margin-bottom: 2.75rem;", // 44px
        
        "mt-11": "margin-top: 2.75rem;", // 44px
        "mr-11": "margin-right: 2.75rem;", // 44px
        "mb-11": "margin-bottom: 2.75rem;", // 44px
        "ml-11": "margin-left: 2.75rem;", // 44px
        "m-12": "margin: 3rem;", // 48px
        "mx-12": "margin-left: 3rem; margin-right: 3rem;", // 48px
        "my-12": "margin-top: 3rem; margin-bottom: 3rem;", // 48px
       
        "mt-12": "margin-top: 3rem;", // 48px
        "mr-12": "margin-right: 3rem;", // 48px
        "mb-12": "margin-bottom: 3rem;", // 48px
        "ml-12": "margin-left: 3rem;", // 48px
        "m-14": "margin: 3.5rem;", // 56px
        "mx-14": "margin-left: 3.5rem; margin-right: 3.5rem;", // 56px
        "my-14": "margin-top: 3.5rem; margin-bottom: 3.5rem;", // 56px
       
        "mt-14": "margin-top: 3.5rem;", // 56px
        "mr-14": "margin-right: 3.5rem;", // 56px
        "mb-14": "margin-bottom: 3.5rem;", // 56px
        "ml-14": "margin-left: 3.5rem;", // 56px
        "m-16": "margin: 4rem;", // 64px
        "mx-16": "margin-left: 4rem; margin-right: 4rem;", // 64px
        "my-16": "margin-top: 4rem; margin-bottom: 4rem;", // 64px
      
        "mt-16": "margin-top: 4rem;", // 64px
        "mr-16": "margin-right: 4rem;", // 64px
        "mb-16": "margin-bottom: 4rem;", // 64px
        "ml-16": "margin-left: 4rem;", // 64px
        "m-20": "margin: 5rem;", // 80px
        "mx-20": "margin-left: 5rem; margin-right: 5rem;", // 80px
        "my-20": "margin-top: 5rem; margin-bottom: 5rem;", // 80px
       
        "mt-20": "margin-top: 5rem;", // 80px
        "mr-20": "margin-right: 5rem;", // 80px
        "mb-20": "margin-bottom: 5rem;", // 80px
        "ml-20": "margin-left: 5rem;", // 80px
        "m-24": "margin: 6rem;", // 96px
        "mx-24": "margin-left: 6rem; margin-right: 6rem;", // 96px
        "my-24": "margin-top: 6rem; margin-bottom: 6rem;", // 96px
       
        "mt-24": "margin-top: 6rem;", // 96px
        "mr-24": "margin-right: 6rem;", // 96px
        "mb-24": "margin-bottom: 6rem;", // 96px
        "ml-24": "margin-left: 6rem;", // 96px
        "m-28": "margin: 7rem;", // 112px
        "mx-28": "margin-left: 7rem; margin-right: 7rem;", // 112px
        "my-28": "margin-top: 7rem; margin-bottom: 7rem;", // 112px
    
        "mt-28": "margin-top: 7rem;", // 112px
        "mr-28": "margin-right: 7rem;", // 112px
        "mb-28": "margin-bottom: 7rem;", // 112px
        "ml-28": "margin-left: 7rem;", // 112px
        "m-32": "margin: 8rem;", // 128px
        "mx-32": "margin-left: 8rem; margin-right: 8rem;", // 128px
        "my-32": "margin-top: 8rem; margin-bottom: 8rem;", // 128px
      
        "mt-32": "margin-top: 8rem;", // 128px
        "mr-32": "margin-right: 8rem;", // 128px
        "mb-32": "margin-bottom: 8rem;", // 128px
        "ml-32": "margin-left: 8rem;", // 128px
        "m-36": "margin: 9rem;", // 144px
        "mx-36": "margin-left: 9rem; margin-right: 9rem;", // 144px
        "my-36": "margin-top: 9rem; margin-bottom: 9rem;", // 144px
       
        "mt-36": "margin-top: 9rem;", // 144px
        "mr-36": "margin-right: 9rem;", // 144px
        "mb-36": "margin-bottom: 9rem;", // 144px
        "ml-36": "margin-left: 9rem;", // 144px
        "m-40": "margin: 10rem;", // 160px
        "mx-40": "margin-left: 10rem; margin-right: 10rem;", // 160px
        "my-40": "margin-top: 10rem; margin-bottom: 10rem;", // 160px
       
        "mt-40": "margin-top: 10rem;", // 160px
        "mr-40": "margin-right: 10rem;", // 160px
        "mb-40": "margin-bottom: 10rem;", // 160px
        "ml-40": "margin-left: 10rem;", // 160px
        "m-44": "margin: 11rem;", // 176px
        "mx-44": "margin-left: 11rem; margin-right: 11rem;", // 176px
        "my-44": "margin-top: 11rem; margin-bottom: 11rem;", // 176px
       
        "mt-44": "margin-top: 11rem;", // 176px
        "mr-44": "margin-right: 11rem;", // 176px
        "mb-44": "margin-bottom: 11rem;", // 176px
        "ml-44": "margin-left: 11rem;", // 176px
        "m-48": "margin: 12rem;", // 192px
        "mx-48": "margin-left: 12rem; margin-right: 12rem;", // 192px
        "my-48": "margin-top: 12rem; margin-bottom: 12rem;", // 192px
        
        "mt-48": "margin-top: 12rem;", // 192px
        "mr-48": "margin-right: 12rem;", // 192px
        "mb-48": "margin-bottom: 12rem;", // 192px
        "ml-48": "margin-left: 12rem;", // 192px
        "m-52": "margin: 13rem;", // 208px
        "mx-52": "margin-left: 13rem; margin-right: 13rem;", // 208px
        "my-52": "margin-top: 13rem; margin-bottom: 13rem;", // 208px
       
        "mt-52": "margin-top: 13rem;", // 208px
        "mr-52": "margin-right: 13rem;", // 208px
        "mb-52": "margin-bottom: 13rem;", // 208px
        "ml-52": "margin-left: 13rem;", // 208px
        "m-56": "margin: 14rem;", // 224px
        "mx-56": "margin-left: 14rem; margin-right: 14rem;", // 224px
        "my-56": "margin-top: 14rem; margin-bottom: 14rem;", // 224px
       
        "mt-56": "margin-top: 14rem;", // 224px
        "mr-56": "margin-right: 14rem;", // 224px
        "mb-56": "margin-bottom: 14rem;", // 224px
        "ml-56": "margin-left: 14rem;", // 224px
        "m-60": "margin: 15rem;", // 240px
        "mx-60": "margin-left: 15rem; margin-right: 15rem;", // 240px
        "my-60": "margin-top: 15rem; margin-bottom: 15rem;", // 240px
        
        "mt-60": "margin-top: 15rem;", // 240px
        "mr-60": "margin-right: 15rem;", // 240px
        "mb-60": "margin-bottom: 15rem;", // 240px
        "ml-60": "margin-left: 15rem;", // 240px
        "m-64": "margin: 16rem;", // 256px
        "mx-64": "margin-left: 16rem; margin-right: 16rem;", // 256px
        "my-64": "margin-top: 16rem; margin-bottom: 16rem;", // 256px
        
        "mt-64": "margin-top: 16rem;", // 256px
        "mr-64": "margin-right: 16rem;", // 256px
        "mb-64": "margin-bottom: 16rem;", // 256px
        "ml-64": "margin-left: 16rem;", // 256px
        "m-72": "margin: 18rem;", // 288px
        "mx-72": "margin-left: 18rem; margin-right: 18rem;", // 288px
        "my-72": "margin-top: 18rem; margin-bottom: 18rem;", // 288px
        
        "mt-72": "margin-top: 18rem;", // 288px
        "mr-72": "margin-right: 18rem;", // 288px
        "mb-72": "margin-bottom: 18rem;", // 288px
        "ml-72": "margin-left: 18rem;", // 288px
        "m-80": "margin: 20rem;", // 320px
        "mx-80": "margin-left: 20rem; margin-right: 20rem;", // 320px
        "my-80": "margin-top: 20rem; margin-bottom: 20rem;", // 320px
       
        "mt-80": "margin-top: 20rem;", // 320px
        "mr-80": "margin-right: 20rem;", // 320px
        "mb-80": "margin-bottom: 20rem;", // 320px
        "ml-80": "margin-left: 20rem;", // 320px
        "m-96": "margin: 24rem;", // 384px
        "mx-96": "margin-left: 24rem; margin-right: 24rem;", // 384px
        "my-96": "margin-top: 24rem; margin-bottom: 24rem;", // 384px
       
        "mt-96": "margin-top: 24rem;", // 384px
        "mr-96": "margin-right: 24rem;", // 384px
        "mb-96": "margin-bottom: 24rem;", // 384px
        "ml-96": "margin-left: 24rem;", // 384px
        "m-auto": "margin: auto;",
        "mx-auto": "margin-left: auto; margin-right: auto;",
        "my-auto": "margin-top: auto; margin-bottom: auto;",
     
        "mt-auto": "margin-top: auto;",
        "mr-auto": "margin-right: auto;",
        "mb-auto": "margin-bottom: auto;",
        "ml-auto": "margin-left: auto;",

        // Space Between Items
        "space-x-0 > * + *": "margin-left: 0px;",
        "space-y-0 > * + *": "margin-top: 0px;",
        "space-x-0.5 > * + *": "margin-left: 0.125rem;", // 2px
        "space-y-0.5 > * + *": "margin-top: 0.125rem;", // 2px
        "space-x-1 > * + *": "margin-left: 0.25rem;", // 4px
        "space-y-1 > * + *": "margin-top: 0.25rem;", // 4px
        "space-x-1.5 > * + *": "margin-left: 0.375rem;", // 6px
        "space-y-1.5 > * + *": "margin-top: 0.375rem;", // 6px
        "space-x-2 > * + *": "margin-left: 0.5rem;", // 8px
        "space-y-2 > * + *": "margin-top: 0.5rem;", // 8px
        "space-x-2.5 > * + *": "margin-left: 0.625rem;", // 10px
        "space-y-2.5 > * + *": "margin-top: 0.625rem;", // 10px
        "space-x-3 > * + *": "margin-left: 0.75rem;", // 12px
        "space-y-3 > * + *": "margin-top: 0.75rem;", // 12px
        "space-x-3.5 > * + *": "margin-left: 0.875rem;", // 14px
        "space-y-3.5 > * + *": "margin-top: 0.875rem;", // 14px
        "space-x-4 > * + *": "margin-left: 1rem;", // 16px
        "space-y-4 > * + *": "margin-top: 1rem;", // 16px
        "space-x-5 > * + *": "margin-left: 1.25rem;", // 20px
        "space-y-5 > * + *": "margin-top: 1.25rem;", // 20px
        "space-x-6 > * + *": "margin-left: 1.5rem;", // 24px
        "space-y-6 > * + *": "margin-top: 1.5rem;", // 24px
        "space-x-7 > * + *": "margin-left: 1.75rem;", // 28px
        "space-y-7 > * + *": "margin-top: 1.75rem;", // 28px
        "space-x-8 > * + *": "margin-left: 2rem;", // 32px
        "space-y-8 > * + *": "margin-top: 2rem;", // 32px
        "space-x-9 > * + *": "margin-left: 2.25rem;", // 36px
        "space-y-9 > * + *": "margin-top: 2.25rem;", // 36px
        "space-x-10 > * + *": "margin-left: 2.5rem;", // 40px
        "space-y-10 > * + *": "margin-top: 2.5rem;", // 40px
        "space-x-11 > * + *": "margin-left: 2.75rem;", // 44px
        "space-y-11 > * + *": "margin-top: 2.75rem;", // 44px
        "space-x-12 > * + *": "margin-left: 3rem;", // 48px
        "space-y-12 > * + *": "margin-top: 3rem;", // 48px
        "space-x-14 > * + *": "margin-left: 3.5rem;", // 56px
        "space-y-14 > * + *": "margin-top: 3.5rem;", // 56px
        "space-x-16 > * + *": "margin-left: 4rem;", // 64px
        "space-y-16 > * + *": "margin-top: 4rem;", // 64px
        "space-x-20 > * + *": "margin-left: 5rem;", // 80px
        "space-y-20 > * + *": "margin-top: 5rem;", // 80px
        "space-x-24 > * + *": "margin-left: 6rem;", // 96px
        "space-y-24 > * + *": "margin-top: 6rem;", // 96px
        "space-x-28 > * + *": "margin-left: 7rem;", // 112px
        "space-y-28 > * + *": "margin-top: 7rem;", // 112px
        "space-x-32 > * + *": "margin-left: 8rem;", // 128px
        "space-y-32 > * + *": "margin-top: 8rem;", // 128px
        "space-x-36 > * + *": "margin-left: 9rem;", // 144px
        "space-y-36 > * + *": "margin-top: 9rem;", // 144px
        "space-x-40 > * + *": "margin-left: 10rem;", // 160px
        "space-y-40 > * + *": "margin-top: 10rem;", // 160px
        "space-x-44 > * + *": "margin-left: 11rem;", // 176px
        "space-y-44 > * + *": "margin-top: 11rem;", // 176px
        "space-x-48 > * + *": "margin-left: 12rem;", // 192px
        "space-y-48 > * + *": "margin-top: 12rem;", // 192px
        "space-x-52 > * + *": "margin-left: 13rem;", // 208px
        "space-y-52 > * + *": "margin-top: 13rem;", // 208px
        "space-x-56 > * + *": "margin-left: 14rem;", // 224px
        "space-y-56 > * + *": "margin-top: 14rem;", // 224px
        "space-x-60 > * + *": "margin-left: 15rem;", // 240px
        "space-y-60 > * + *": "margin-top: 15rem;", // 240px
        "space-x-64 > * + *": "margin-left: 16rem;", // 256px
        "space-y-64 > * + *": "margin-top: 16rem;", // 256px
        "space-x-72 > * + *": "margin-left: 18rem;", // 288px
        "space-y-72 > * + *": "margin-top: 18rem;", // 288px
        "space-x-80 > * + *": "margin-left: 20rem;", // 320px
        "space-y-80 > * + *": "margin-top: 20rem;", // 320px
        "space-x-96 > * + *": "margin-left: 24rem;", // 384px
        "space-y-96 > * + *": "margin-top: 24rem;", // 384px
        "space-x-px > * + *": "margin-left: 1px;",
        "space-y-px > * + *": "margin-top: 1px;",
        "space-y-reverse > * + *": "--tw-space-y-reverse: 1;",
        "space-x-reverse > * + *": "--tw-space-x-reverse: 1;",
    }

    for key, value := range spacing {
        cssClasses[key] = value
    }
}