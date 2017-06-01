$(function () {
	$('body').fadeIn();

	// Submit forms
	$('.modal form').submit(function (event) {
		event.preventDefault();
		var data = getFormData($(this));
		$.ajax({
			type: 'POST',
			url: $(this).attr('action'),
			contentType: 'application/json; charset=utf-8',
			dataType: 'json',
			data: JSON.stringify(data),
			success: function (resp) {
				$('#nav-unauthed').hide();
				$('#username').text(resp.data.username);
				$('#nav-authed').fadeIn();
			},
			error: function (resp) {
				console.log(resp);
			}
		});
		$(this).parents('.modal').modal('toggle')
	})

	// Dismiss authenticated view
	$('#nav-authed button').click(function () {
		$('#nav-authed').hide();
		$('#username').text('');
		$('#nav-unauthed').fadeIn();
	})
})

function getFormData($form){
	var dataArray = $form.serializeArray();
	var dataObject = {};

	$.map(dataArray, function(n, i){
		dataObject[n.name] = n.value;
	});
	return dataObject;
}
