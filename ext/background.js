chrome.contextMenus.create({
	"title": chrome.i18n.getMessage("findImage"),
	"contexts": ["image"],
	"onclick": function(info, tab) {
		chrome.tabs.getSelected(null, function(tab){
			let index = tab.index + 1,
				url = "http://127.0.0.1/zapr?url=" + encodeURIComponent(info.srcUrl);
				fetch(url,{mode:'no-cors'})
					.then(function(){chrome.tabs.create({url: "http://127.0.0.1/img", index: index})});
		});
	}
});