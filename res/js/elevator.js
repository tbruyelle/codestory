$(document).ready(function() {
	window.setInterval(function(event) {
		$.pjax({url:'/', container:'#pjax-container'})
	}, 5000)
})

