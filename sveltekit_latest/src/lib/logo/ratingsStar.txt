<script lang="ts">
	// your script goes here

</script>

<!-- markup (zero or more items) goes here -->

<div id="half-stars-example">
	<div class="rating-group">
		<input
			class="rating__input rating__input--none"
			checked
			name="rating2"
			id="rating2-0"
			value="0"
			type="radio"
		/>
		<label aria-label="0 stars" class="rating__label" for="rating2-0">&nbsp;</label>
		<label aria-label="0.5 stars" class="rating__label rating__label--half" for="rating2-05"
			><i class="rating__icon rating__icon--star fa fa-star-half" /></label
		>
		<input  class="rating__input" name="rating2" id="rating2-05" value="0.5"  type="radio" />
		<label aria-label="1 star" class="rating__label" for="rating2-10"
			><i class="rating__icon rating__icon--star fa fa-star" /></label
		>
		<input class="rating__input" name="rating2" id="rating2-10" value="1" type="radio" />
		<label aria-label="1.5 stars" class="rating__label rating__label--half" for="rating2-15"
			><i class="rating__icon rating__icon--star fa fa-star-half" /></label
		>
		<input class="rating__input" name="rating2" id="rating2-15" value="1.5" type="radio" />
		<label aria-label="2 stars" class="rating__label" for="rating2-20"
			><i class="rating__icon rating__icon--star fa fa-star" /></label
		>
		<input class="rating__input" name="rating2" id="rating2-20" value="2" type="radio" />
		<label aria-label="2.5 stars" class="rating__label rating__label--half" for="rating2-25"
			><i class="rating__icon rating__icon--star fa fa-star-half" /></label
		>
		<input class="rating__input" name="rating2" id="rating2-25" value="2.5" type="radio" checked />
		<label aria-label="3 stars" class="rating__label" for="rating2-30"
			><i class="rating__icon rating__icon--star fa fa-star" /></label
		>
		<input class="rating__input" name="rating2" id="rating2-30" value="3" type="radio" />
		<label aria-label="3.5 stars" class="rating__label rating__label--half" for="rating2-35"
			><i class="rating__icon rating__icon--star fa fa-star-half" /></label
		>
		<input class="rating__input" name="rating2" id="rating2-35" value="3.5" type="radio" />
		<label aria-label="4 stars" class="rating__label" for="rating2-40"
			><i class="rating__icon rating__icon--star fa fa-star" /></label
		>
		<input class="rating__input" name="rating2" id="rating2-40" value="4" type="radio" />
		<label aria-label="4.5 stars" class="rating__label rating__label--half" for="rating2-45"
			><i class="rating__icon rating__icon--star fa fa-star-half" /></label
		>
		<input class="rating__input" name="rating2" id="rating2-45" value="4.5" type="radio" />
		<label aria-label="5 stars" class="rating__label" for="rating2-50"
			><i class="rating__icon rating__icon--star fa fa-star" /></label
		>
		<input class="rating__input" name="rating2" id="rating2-50" value="5" type="radio" />
	</div>
</div>

<style lang="css">


	

	#half-stars-example .rating-group {
		display: inline-flex;
	}
	#half-stars-example .rating__icon {
		pointer-events: none;
	}
	#half-stars-example .rating__input {
		position: absolute !important;
		left: -9999px !important;
	}
	#half-stars-example .rating__label {
		cursor: pointer;
		/* if you change the left/right padding, update the margin-right property of .rating__label--half as well. */
		padding: 0 0.1em;
		font-size: 2rem;
	}
	#half-stars-example .rating__label--half {
		padding-right: 0;
		margin-right: -0.6em;
		z-index: 2;
	}
	#half-stars-example .rating__icon--star {
		color: orange;
	}
	#half-stars-example .rating__icon--none {
		color: #eee;
	}
	#half-stars-example .rating__input--none:checked + .rating__label .rating__icon--none {
		color: red;
	}
	#half-stars-example .rating__input:checked ~ .rating__label .rating__icon--star {
		color: #ddd;
	}
	#half-stars-example .rating-group:hover .rating__label .rating__icon--star,
	#half-stars-example .rating-group:hover .rating__label--half .rating__icon--star {
		color: orange;
	}
	#half-stars-example .rating__input:hover ~ .rating__label .rating__icon--star,
	#half-stars-example .rating__input:hover ~ .rating__label--half .rating__icon--star {
		color: #ddd;
	}
	#half-stars-example
		.rating-group:hover
		.rating__input--none:not(:hover)
		+ .rating__label
		.rating__icon--none {
		color: #eee;
	}
	#half-stars-example .rating__input--none:hover + .rating__label .rating__icon--none {
		color: red;
	}

	
</style>
