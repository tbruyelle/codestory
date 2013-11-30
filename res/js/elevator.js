$(document).ready(function() {
	$('#refresh').on('click', function(event) {
		$.pjax({url:'/', container:'#pjax-container'})
	})
})

