function getIndex(hash, callback) {
	if( typeof hash === "string" ) {
		node.files.cat(hash + "/index.json", function(err, data) {
			if(err) {
				return console.error('Error - ipfs files cat', err, res)
			}

			if( typeof callback === "function" ) {
				callback(data)
			}
	})
	}
}
