<script>
    let { class: additionalClasses = "", onclick=null, children, ...restProps } = $props();
</script>

<button 
    {...restProps} 
    {onclick}
    class="bg-blue-500 hover:bg-blue-600 font-medium text-white rounded-md px-4 {additionalClasses}"
>
    {@render children?.()}
</button>