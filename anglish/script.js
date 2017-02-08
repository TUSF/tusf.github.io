var $container = document.getElementById('container');
var $backdrop = document.getElementById('backdrop');
var $highlights = document.getElementById('highlights');
var $textarea = document.getElementById('text');

// yeah, browser sniffing sucks, but there are browser-specific quirks to handle that are not a matter of feature detection
var ua = window.navigator.userAgent.toLowerCase();
var isIE = !!ua.match(/msie|trident\/7|edge/);
var isWinPhone = ua.indexOf('windows phone') !== -1;
var isIOS = !isWinPhone && !!ua.match(/ipad|iphone|ipod/);


function applyHighlights(text) {
	text = text
		.replace(/\n$/g, '\n\n')
		.replace(OldFrenchTerms, '<mark class="oldfrench">$&</mark>')
		.replace(AngloNormanTerms, '<mark class="anglonorman">$&</mark>');
	if (isIE) {
		// IE wraps whitespace differently in a div vs textarea, this fixes it
		text = text.replace(/ /g, ' <wbr>');
	}
	
	return text;
}

function handleInput() {
	$highlights.innerHTML = applyHighlights($textarea.value);
	handleScroll()
}

function handleScroll() {
	$backdrop.scrollTop = $textarea.scrollTop;
	$backdrop.scrollLeft = $textarea.scrollLeft;
}

function fixIOS() {
	// iOS adds 3px of (unremovable) padding to the left and right of a textarea, so adjust highlights div to match
	$highlights.style.paddingLeft = "3px"
	$highlights.style.paddingRight = "3px"
}

function bindEvents() {
	//Timeout, because "keyup" takes too long, and "keydown" fires BEFORE the letter is added.
	//Same for cut and paste.
	$textarea.addEventListener('keydown', function(){setTimeout(handleInput,0)});
	$textarea.addEventListener('cut', function(){setTimeout(handleInput,0)});
	$textarea.addEventListener('paste', function(){setTimeout(handleInput,0)});
	$textarea.addEventListener('change', function(){setTimeout(handleInput,0)});
	$textarea.addEventListener('scroll', handleScroll);
}

if (isIOS) {
	fixIOS();
}

bindEvents();
handleInput();

// Retreive terms.

var AngloNormanTerms;
var OldFrenchTerms;
(function() {
	//Anglo-Norman
		anRequest = new XMLHttpRequest();
		anRequest.onreadystatechange = function(){
		if (anRequest.readyState === XMLHttpRequest.DONE) {
			if (anRequest.status === 200) {
				terms = anRequest.responseText.split("\n")
				reg = "\\b(" + terms.join("|") + ")\\b"
				AngloNormanTerms = new RegExp(reg, 'ig')
				//Super long regex, because I'm a bad coder.
				//If you have a better way, comit it, pls.
			}
		}
		}
		anRequest.open('GET', 'anglonorman.txt', true);
		anRequest.send(null);
	
	//Old French
		ofRequest = new XMLHttpRequest();
		ofRequest.onreadystatechange = function(){
		if (ofRequest.readyState === XMLHttpRequest.DONE) {
			if (ofRequest.status === 200) {
				terms = anRequest.responseText.split("\n")
				reg = "\\b(" + terms.join("|") + ")\\b"
				AngloNormanTerms = new RegExp(reg, 'ig')
				//Super long regex, because I'm a bad coder.
				//If you have a better way, comit it, pls.
			}
		}
		}
		ofRequest.open('GET', 'anglonorman.txt', true);
		ofRequest.send(null);
})();
