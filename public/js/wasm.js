window.onload = () => {
	if (!WebAssembly.instantiateStreaming) {
		WebAssembly.instantiateStreaming = async (resp, importObject) => {
			const source = await (await resp).arrayBuffer();
			return await WebAssembly.instantiate(source, importObject);
		};
	}

	const go = new Go();
	WebAssembly.instantiateStreaming(fetch("kuronan-dash.wasm"), go.importObject)
	.then((result) => {
		document.getElementById("loading").remove();
		go.run(result.instance);
	})
	.catch((err) => {
		console.error(err);
	});
}