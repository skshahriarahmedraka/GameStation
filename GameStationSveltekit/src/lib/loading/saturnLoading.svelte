<script>
	// your script goes here
</script>

<!-- markup (zero or more items) goes here -->

<!-- about -->
<!-- <div class="about h-screen w-full flex justify-center items-center ?">
    <a class="bg_links social portfolio" href="https://www.rafaelalucas.com" target="_blank">
       <span class="icon"></span>
    </a>
    <a class="bg_links social dribbble" href="https://dribbble.com/rafaelalucas" target="_blank">
       <span class="icon"></span>
    </a>
    <a class="bg_links social linkedin" href="https://www.linkedin.com/in/rafaelalucas/" target="_blank">
       <span class="icon"></span>
    </a>
    <a class="bg_links logo"></a>
 </div> -->
<!-- end about -->
<!-- bg-[#3c4359]  -->
<div class=" {$$props.class} flex h-screen w-full items-center justify-center  bg-[#121212]  ">
	<div class="content ">
		<div class="planet">
			<div class="ring" />
			<div class="cover-ring" />
			<div class="spots">
				<span />
				<span />
				<span />
				<span />
				<span />
				<span />
				<span />
			</div>
		</div>
		<p class="  ">loading</p>
	</div>
</div>

<style>
	@import url('https://fonts.googleapis.com/css?family=Roboto+Mono&display=swap');

	/* 
body {
  margin: 0;
  padding: 0;
  display: flex;
  align-items: center;
  justify-content: center;
  height: 100vh;
  width: 100%;
  background-color: #3c4359;
} */

	.content {
		width: 300px;
		height: 300px;
		position: relative;
		display: flex;
		justify-content: center;
		align-items: center;
	}
	.content .planet {
		width: 65%;
		height: 65%;
		background-color: #546c8c;
		border-radius: 100%;
		position: absolute;
		display: flex;
		align-items: center;
		transform-origin: center center;
		box-shadow: inset 2px -10px 0px rgba(0, 0, 0, 0.1);
		animation: planet 5s ease infinite alternate;
		/* planet ring */
		/* to cover the back of the ring */
		/* planet spots */
	}
	@keyframes planet {
		0% {
			transform: rotate(10deg);
		}
		100% {
			transform: rotate(-10deg);
		}
	}
	.content .planet .ring {
		position: absolute;
		width: 300px;
		height: 300px;
		border-radius: 100%;
		background-color: #bacbd9;
		display: flex;
		align-items: center;
		justify-content: center;
		transform-origin: 33% center;
		box-shadow: 2px -10px 0px rgba(0, 0, 0, 0.1), inset -5px -10px 0px rgba(0, 0, 0, 0.1);
		animation: ring 3s ease infinite;
		/* small ball */
		/* inner ring */
	}
	@keyframes ring {
		0% {
			transform: rotateX(110deg) rotateZ(0deg) translate(-50px, 5px);
		}
		100% {
			transform: rotateX(110deg) rotateZ(360deg) translate(-50px, 5px);
		}
	}
	.content .planet .ring:before {
		content: '';
		position: absolute;
		width: 10px;
		height: 30px;
		border-radius: 100%;
		background-color: #7ea1bf;
		z-index: 2;
		left: calc(0px - 5px);
		box-shadow: inset -3px 3px 0px rgba(0, 0, 0, 0.2);
	}
	.content .planet .ring:after {
		content: '';
		position: absolute;
		width: 240px;
		height: 240px;
		border-radius: 100%;
		background-color: #7ea1bf;
		box-shadow: inset 2px -10px 0px rgba(0, 0, 0, 0.1);
	}
	.content .planet .cover-ring {
		position: absolute;
		width: 100%;
		height: 50%;
		border-bottom-left-radius: 80%;
		border-bottom-right-radius: 80%;
		border-top-left-radius: 100px;
		border-top-right-radius: 100px;
		transform: translate(0px, -17px);
		background-color: #546c8c;
		z-index: 2;
		box-shadow: inset 0px -2px 0px rgba(0, 0, 0, 0.1);
	}
	.content .planet .spots {
		width: 100%;
		height: 100%;
		display: flex;
		align-items: center;
		justify-content: center;
		position: absolute;
		z-index: 2;
	}
	.content .planet .spots span {
		width: 30px;
		height: 30px;
		background-color: #3c4359;
		position: absolute;
		border-radius: 100%;
		box-shadow: inset -2px 3px 0px rgba(0, 0, 0, 0.3);
		animation: dots 5s ease infinite alternate;
	}
	@keyframes dots {
		0% {
			box-shadow: inset -3px 3px 0px rgba(0, 0, 0, 0.3);
		}
		100% {
			box-shadow: inset 3px 3px 0px rgba(0, 0, 0, 0.3);
		}
	}
	.content .planet .spots span:nth-child(1) {
		top: 20px;
		right: 50px;
	}
	.content .planet .spots span:nth-child(2) {
		top: 40px;
		left: 50px;
		width: 15px;
		height: 15px;
	}
	.content .planet .spots span:nth-child(3) {
		top: 80px;
		left: 20px;
		width: 25px;
		height: 25px;
	}
	.content .planet .spots span:nth-child(4) {
		top: 80px;
		left: 90px;
		width: 40px;
		height: 40px;
	}
	.content .planet .spots span:nth-child(5) {
		top: 160px;
		left: 70px;
		width: 15px;
		height: 15px;
	}
	.content .planet .spots span:nth-child(6) {
		top: 165px;
		left: 125px;
		width: 10px;
		height: 10px;
	}
	.content .planet .spots span:nth-child(7) {
		top: 90px;
		left: 150px;
		width: 15px;
		height: 15px;
	}
	.content p {
		color: #bacbd9;
		font-size: 20px;
		z-index: 2;
		position: absolute;
		bottom: -20px;
		font-family: 'Roboto Mono', monospace;
		animation: text 4s ease infinite;
		width: 100px;
		text-align: center;
	}
	@keyframes text {
		0% {
			transform: translateX(-30px);
			letter-spacing: 0px;
			color: #bacbd9;
		}
		25% {
			letter-spacing: 3px;
			color: #7ea1bf;
		}
		50% {
			transform: translateX(30px);
			letter-spacing: 0px;
			color: #bacbd9;
		}
		75% {
			letter-spacing: 3px;
			color: #7ea1bf;
		}
		100% {
			transform: translateX(-30px);
			letter-spacing: 0px;
			color: #bacbd9;
		}
	}
</style>
