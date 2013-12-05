$(document).ready(function() {
	$.pjax.defaults.scrollTo = false;
	window.setInterval(function(event) {
		$.pjax({url:'/', container:'#pjax-container'})
	}, 2000)
})

