$(function () {
	$("body").fadeIn();

	// Submit forms
	$(".modal button[type='submit']").click(function () {
		var form = $(this).parents(".modal").find("form");
		var data = getFormData(form);
		$.ajax({
           type: "POST",
           url: form.attr("action"),
		   contentType: "application/json; charset=utf-8",
		   dataType: "json",
           data: JSON.stringify(data),
           success: function (resp) {
			   $("#nav-unauthed").hide();
			   $("#username").text(resp.data.username);
			   $("#nav-authed").fadeIn();
               console.log(resp);
           },
		   error: function (resp) {
			   console.log(resp);
		   }
       });
	})

	// Dismiss authenticated view
	$("#nav-authed button").click(function () {
		$("#nav-authed").hide();
		$("#username").text("");
		$("#nav-unauthed").fadeIn();
	})
})

function getFormData($form){
    var dataArray = $form.serializeArray();
    var dataObject = {};

    $.map(dataArray, function(n, i){
        dataObject[n['name']] = n['value'];
    });
    return dataObject;
}
