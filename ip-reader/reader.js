function getIndex(hash, callback) {
	if( typeof hash === "string" && typeof callback === "function" ) {
		node.files.cat(hash + "/index.json", function(err, data) {
			if(err) {
				return console.error('Error - ipfs files cat', err)
			}
			index = JSON.parse(data)

			node.files.cat(hash + "/" + index.cover, function(err, idata) {
				if(err) {
					return console.error('Error - ipfs files cat', err)
				}
				mediatype = "image/"
				switch ( index.cover.substr(index.cover.lastIndexOf('.') + 1) ) {
					case 'png':
						mediatype += "png"
					case 'gif':
						mediatype += "gif"
					case 'jpg':
					case 'jpeg':
						mediatype += "jpeg"
				}

				index.coverdata = "data:" + mediatype + ";base64," + idata.toString('base64')
				callback(index)
			})
		})
	}
}
