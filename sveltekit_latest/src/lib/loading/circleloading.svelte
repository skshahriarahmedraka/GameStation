<div class="wrap">
	<div class="loaders">
		<svg height="40" width="40" class="loader loader-1">
			<g class="cir">
				<path d="m 20 5 a 1 1 0 0 0 0 30 a 1 1 0 0 0 0 -30" stroke="rgba(0,0,0,0.2)" stroke-width="3" fill="none" stroke-linecap="round"/>
			</g>
			<g class="go">
				<path d="m 20 5 c 0 0 15 0 15 15" stroke="#ff5454" stroke-width="3" fill="none" stroke-linecap="round"/>
			</g>
		</svg>
		<svg height="40" width="40" class="loader loader-2">
			<g class="cir">
				<path d="m 20 5 a 1 1 0 0 0 0 30 a 1 1 0 0 0 0 -30" stroke="rgba(0,0,0,0.2)" stroke-width="3" fill="none" stroke-linecap="round"/>
			</g>
			<g class="go">
				<path d="m 20 5 a 1 1 0 0 0 0 30" stroke="#ffbb54" stroke-width="3" fill="none" stroke-linecap="round"/>
			</g>
		</svg>
		<svg height="40" width="40" class="loader loader-3">
			<g class="cir">
				<path d="m 20 5 a 1 1 0 0 0 0 30 a 1 1 0 0 0 0 -30" stroke="rgba(0,0,0,0.2)" stroke-width="3" fill="none" stroke-linecap="round"/>
			</g>
			<g class="go">
				<path d="m 20 35 a 1 1 0 0 1 0 -30 c 0 0 15 0 15 15" stroke="#54ff5a" stroke-width="3" fill="none" stroke-linecap="round"/>
			</g>
		</svg>
		<svg height="40" width="40" class="loader loader-4">
			<g class="load-4-1">
				<path d="m 20 5 a 1 1 0 0 1 0 30 c 0 0 -15 0 -15 -15" stroke="#4089ff" stroke-width="3" fill="none" stroke-linecap="round"/>
			</g>
			<g class="load-4-2">
				<path d="m 20 10 a 1 1 0 0 1 0 20" stroke="#4089ff" stroke-width="3" fill="none" stroke-linecap="round"/>
			</g>
			<g class="load-4-3">
				<path d="m 20 15 c 0 0 5 0 5 5" stroke="#4089ff" stroke-width="3" fill="none" stroke-linecap="round"/>
			</g>
		</svg>
		<svg height="40" width="40" class="loader loader-5">
			<g class="cir">
				<path d="m 20 5 a 1 1 0 0 0 0 30 a 1 1 0 0 0 0 -30" stroke="rgba(0,0,0,0.2)" stroke-width="3" fill="none" stroke-linecap="round"/>
			</g>
			<g class="go">
				<path d="m 20 5 a 1 1 0 0 0 0 30 a 1 1 0 0 0 0 -30" stroke="#5457ff" stroke-width="3" fill="none" stroke-linecap="round"/>
			</g>
		</svg>
	</div>
</div>

<style>
 
.wrap {
	width: fit-content;
	height: 60px;
	background: #333;
	border-radius: 10px;
	padding: 10px;
	position: absolute;
	top: 50%;
	left: 50%;
	transform: translate(-50%,-50%);
	transition: 0.3s;
	cursor: pointer;
  box-shadow: 2px 2px 5px rgba(0,0,0,0.3);
}
.loaders svg {
	padding: 10px;
}
svg {
	display: inline-block;
	transition: 0.3s;
	border-radius: 5px;
}
.go path {
	animation-name: ro;
	animation-duration: 1s;
	animation-timing-function: linear;
	animation-iteration-count: infinite;
	transform-origin: 50% 50%;
}

@keyframes ro {
	100% {
		transform: rotate(360deg) translate(0,0);
	}
}

/* flip the path rather than using negative 'stroke-dashoffset' */
/* otherwise, it will not work correctly on some browsers. */
.loader-5 .go path {
	stroke-dasharray: 100;
	stroke-dashoffset: 100;
	transform: scaleX(-1);
	transform-origin: 50% 50%;
}

.loader-5 .go path {
	animation-name: ro-5;
}

@keyframes ro-5 {
	100% {
		stroke-dasharray: 200;
		stroke-dashoffset: 0;
	}
}


.loader-4 {
	filter: drop-shadow(3px 3px 0px rgba(0,0,0,0.2));
}
.loader-4 .load-4-1, .loader-4 .load-4-2, .loader-4 .load-4-3 {
	transform-origin: 50% 50%;
	animation-iteration-count: infinite;
	
}
.loader-4 .load-4-1 path, .loader-4 .load-4-2 path, .loader-4 .load-4-3 path {
	box-shadow: 2px 2px 0 #fff inset;
}
.loader-4 .load-4-1 {
	animation-name: ro-4-1;
	animation-duration: 1s;
}
.loader-4 .load-4-2 {
	animation-name: ro-4-2;
	animation-duration: 1s;
	animation-delay: 0.5s;
}
.loader-4 .load-4-3 {
	animation-name: ro-4-3;
	animation-duration: 1s;
	animation-delay: 1s;
}


@keyframes ro-4-1 {
	100% {
		transform: rotate(360deg) translate(0,0);
	}
}
@keyframes ro-4-2 {
	100% {
		transform: rotate(360deg) translate(0,0);
	}
}
@keyframes ro-4-3 {
	100% {
		transform: rotate(360deg) translate(0,0);
	}
}
</style>