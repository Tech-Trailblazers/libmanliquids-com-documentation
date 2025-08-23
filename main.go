package main // Define the main package

import (
	"bytes"         // Provides bytes buffer and manipulation utilities
	"io"            // Provides I/O primitives like Reader and Writer
	"log"           // Provides logging functionalities
	"net/http"      // Provides HTTP client and server implementations
	"net/url"       // Provides URL parsing and encoding utilities
	"os"            // Provides file system and OS-level utilities
	"path/filepath" // Provides utilities for file path manipulation
	"regexp"        // Provides support for regular expressions
	"strings"       // Provides string manipulation utilities
	"time"          // Provides time-related functions

	"golang.org/x/net/html" // Provides support for parsing HTML documents
)

func main() {
	remoteAPIURL := []string{
		"https://libmanliquids.com/products/freedom-concentrated-hardwood-cleaner",
		"https://libmanliquids.com/products/freedom-concentrated-multi-surface-floor-cleaner",
		"https://libmanliquids.com/products/2-sided-microfiber-mop",
		"https://libmanliquids.com/products/4-gallon-clean-rinse-bucket",
		"https://libmanliquids.com/products/4-gallon-clean-rinse-bucket-with-wringer",
		"https://libmanliquids.com/products/all-purpose-bucket",
		"https://libmanliquids.com/products/clean-rinse-bucket",
		"https://libmanliquids.com/products/freedom-floor-duster",
		"https://libmanliquids.com/products/freedom-floor-duster-refill",
		"https://libmanliquids.com/products/freedom-spray-mop",
		"https://libmanliquids.com/products/freedom-spray-mop-refill",
		"https://libmanliquids.com/products/gator-mop",
		"https://libmanliquids.com/products/gator-mop-refill",
		"https://libmanliquids.com/products/deluxe-cleaning-caddy",
		"https://libmanliquids.com/products/microfiber-dust-mop",
		"https://libmanliquids.com/products/microfiber-dust-mop-refill",
		"https://libmanliquids.com/products/microfiber-wet-dry-floor-mop",
		"https://libmanliquids.com/products/nitty-gritty-all-surface-roller-mop",
		"https://libmanliquids.com/products/nitty-gritty-roller-mop-refill",
		"https://libmanliquids.com/products/easy-roller-mop",
		"https://libmanliquids.com/products/easy-roller-mop-refill",
		"https://libmanliquids.com/products/scrubster-mop",
		"https://libmanliquids.com/products/swivel-duster",
		"https://libmanliquids.com/products/tornado-twist-mop",
		"https://libmanliquids.com/products/tornado-mop-refill",
		"https://libmanliquids.com/products/255-utility-bucket",
		"https://libmanliquids.com/products/wonder-mop",
		"https://libmanliquids.com/products/wonder-mop-refill",
		"https://libmanliquids.com/products/wood-floor-roller-mop",
		"https://libmanliquids.com/products/wood-floor-roller-mop-refill",
		"https://libmanliquids.com/products/780-all-purpose-wet-mop",
		"https://libmanliquids.com/products/1055-4-gallon-bucket",
		"https://libmanliquids.com/products/1056-4-gallon-bucket-with-wringer",
		"https://libmanliquids.com/products/985-dust-mop-handle",
		"https://libmanliquids.com/products/927-extra-large-microfiber-floor-mop-refill",
		"https://libmanliquids.com/products/934-heavy-duty-bucket-wringer",
		"https://libmanliquids.com/products/781-premium-blue-blend-wet-mop",
		"https://libmanliquids.com/products/979-all-purpose-heavy-duty-wet-mop",
		"https://libmanliquids.com/products/1013-microfiber-all-purpose-cleaning-pad",
		"https://libmanliquids.com/products/1010-microfiber-cleaning-system",
		"https://libmanliquids.com/products/1011-microfiber-fingers-dusting-pad",
		"https://libmanliquids.com/products/1095-one-piece-bucket-wringer",
		"https://libmanliquids.com/products/956-roller-mop-scrub-brush-refill",
		"https://libmanliquids.com/products/1369-wet-floor-sign",
		"https://libmanliquids.com/products/extra-large-precision-angle-broom",
		"https://libmanliquids.com/products/xl-precision-angle-broom-clean-fibers-dustpan",
		"https://libmanliquids.com/products/large-precision-angle-indoor-outdoor-broom",
		"https://libmanliquids.com/products/precision-angle-indoor-outdoor-broom",
		"https://libmanliquids.com/products/shaped-duster-brush",
		"https://libmanliquids.com/products/smooth-sweep-push-broom",
		"https://libmanliquids.com/products/upright-dustpan",
		"https://libmanliquids.com/products/whisk-broom-with-dustpan",
		"https://libmanliquids.com/products/824-18-multi-surface-heavy-duty-push-broom-2",
		"https://libmanliquids.com/products/826-18-fiberforce-rough-surface-push-broom",
		"https://libmanliquids.com/products/800-smooth-surface-push-broom",
		"https://libmanliquids.com/products/1292-24-fiberforce-multi-surface-push-broom",
		"https://libmanliquids.com/products/805-multi-surface-push-broom",
		"https://libmanliquids.com/products/801-smooth-surface-push-broom",
		"https://libmanliquids.com/products/906-dustpan-whisk-broom",
		"https://libmanliquids.com/products/904-indoor-outdoor-angle-broom",
		"https://libmanliquids.com/products/905-indoor-outdoor-angle-broom-dustpan",
		"https://libmanliquids.com/products/581-industrial-grade-dustpan",
		"https://libmanliquids.com/products/1168-large-scoop-dustpan",
		"https://libmanliquids.com/products/915-lobby-broom",
		"https://libmanliquids.com/products/916-closed-lid-dustpan",
		"https://libmanliquids.com/products/918-open-lid-dustpan",
		"https://libmanliquids.com/products/big-feather-duster-6-pack",
		"https://libmanliquids.com/products/gentle-touch-refills",
		"https://libmanliquids.com/products/1341-ultra-absorbent-towels",
		"https://libmanliquids.com/products/dishmatic-non-scratch-scrubber-dish-wand-refills",
		"https://libmanliquids.com/products/dishmatic-general-purpose-dish-wand-refills-3-pack",
		"https://libmanliquids.com/products/dishmatic-i-stand-dish-wand",
		"https://libmanliquids.com/products/bottle-straw-cleaning-kit",
		"https://libmanliquids.com/products/bottle-brush",
		"https://libmanliquids.com/products/palm-brush",
		"https://libmanliquids.com/products/sink-caddy",
		"https://libmanliquids.com/products/tile-grout-brush",
		"https://libmanliquids.com/products/soft-touch-dust-cloth",
		"https://libmanliquids.com/products/small-scrub-brush",
		"https://libmanliquids.com/products/power-scrub-dots-kitchen-dish-wipes",
		"https://libmanliquids.com/products/power-scrub-brush",
		"https://libmanliquids.com/products/pot-pan-scrubbing-dish-wand-with-scrub-brush-refills",
		"https://libmanliquids.com/products/pot-pan-scrubbing-dish-wand-with-scrub-brush",
		"https://libmanliquids.com/products/microfiber-duster",
		"https://libmanliquids.com/products/long-handle-scrub-brush",
		"https://libmanliquids.com/products/kitchen-brush",
		"https://libmanliquids.com/products/heavy-duty-scrub-brush",
		"https://libmanliquids.com/products/hand-nail-brush",
		"https://libmanliquids.com/products/glass-dish-sponge",
		"https://libmanliquids.com/products/glass-dish-wand-refills",
		"https://libmanliquids.com/products/glass-dish-wand-with-scrub-brush",
		"https://libmanliquids.com/products/gentle-touch-foaming-dish-wand",
		"https://libmanliquids.com/products/flexible-microfiber-wand",
		"https://libmanliquids.com/products/everyday-dusting-cloths",
		"https://libmanliquids.com/products/easy-grip-scrub-brush",
		"https://libmanliquids.com/products/dish-brush",
		"https://libmanliquids.com/products/curved-kitchen-brush",
		"https://libmanliquids.com/products/culinary-brush",
		"https://libmanliquids.com/products/brass-pot-brush",
		"https://libmanliquids.com/products/big-job-kitchen-brush",
		"https://libmanliquids.com/products/all-purpose-refills",
		"https://libmanliquids.com/products/all-purpose-scrubbing-dish-wand",
		"https://libmanliquids.com/products/all-purpose-kitchen-brush",
		"https://libmanliquids.com/products/angled-toilet-bowl-brush",
		"https://libmanliquids.com/products/designer-bowl-brush-caddy",
		"https://libmanliquids.com/products/fiberforce-tile-grout-brush",
		"https://libmanliquids.com/products/premium-bowl-brush-and-caddy",
		"https://libmanliquids.com/products/megaforce-premium-plunger-caddy",
		"https://libmanliquids.com/products/megaforce-combo-toilet-brush-plunger",
		"https://libmanliquids.com/products/all-purpose-non-scratch-sponges-3-pack",
		"https://libmanliquids.com/products/all-purpose-odor-resistant-sponges-3-pack",
		"https://libmanliquids.com/products/heavy-duty-easy-rinse-sponges",
		"https://libmanliquids.com/products/non-scratch-easy-rinse-sponge-3-pack",
		"https://libmanliquids.com/products/clean-shine-microfiber-sponge",
		"https://libmanliquids.com/products/copper-power-scrubs",
		"https://libmanliquids.com/products/no-knees-floor-scrub",
		"https://libmanliquids.com/products/scrub-sponges-suction-hanger",
		"https://libmanliquids.com/products/stainless-steel-power-scrubs",
		"https://libmanliquids.com/products/tile-tub-scrub",
		"https://libmanliquids.com/products/tile-tub-scrub-refills",
		"https://libmanliquids.com/products/548-acid-brush",
		"https://libmanliquids.com/products/532-dual-surface-scrub-brush",
		"https://libmanliquids.com/products/547-floor-scrubber",
		"https://libmanliquids.com/products/525-iron-handle-scrub-brush",
		"https://libmanliquids.com/products/521-all-natural-tampico-soft-scrub-brush",
		"https://libmanliquids.com/products/hardwood-floor-polish-and-protector",
		"https://libmanliquids.com/products/multi-surface-everyday-floor-cleaner",
		"https://libmanliquids.com/products/hardwood-floor-everyday-cleaner",
		"https://libmanliquids.com/products/1063-concentrated-window-cleaner",
		"https://libmanliquids.com/products/1064-professional-window-cleaner",
		"https://libmanliquids.com/products/brass-grill-brush",
		"https://libmanliquids.com/products/latex-disposable-gloves-10-pack",
		"https://libmanliquids.com/products/latex-disposable-gloves-50-pack",
		"https://libmanliquids.com/products/nitrile-disposable-gloves-10-pack",
		"https://libmanliquids.com/products/nitrile-disposable-gloves-50-pack",
		"https://libmanliquids.com/products/vent-brush",
		"https://libmanliquids.com/products/586-lambswool-duster",
		"https://libmanliquids.com/products/612-12-foot-extension-handle",
		"https://libmanliquids.com/products/613-16-foot-extension-handle",
		"https://libmanliquids.com/products/611-8-foot-extension-handle",
		"https://libmanliquids.com/products/191-flex-blade-floor-squeegee",
		"https://libmanliquids.com/products/1014-professional-flex-blade-floor-squeegee",
		"https://libmanliquids.com/products/182-all-purpose-squeegee",
		"https://libmanliquids.com/products/194-2-in-1-window-washer",
		"https://libmanliquids.com/products/193-window-glass-washer",
		"https://libmanliquids.com/products/1066-window-cleaning-bucket",
		"https://libmanliquids.com/products/188-window-washer",
		"https://libmanliquids.com/products/1065-window-cleaning-all-one-kit",
		"https://libmanliquids.com/products/precision-angle-indoor-outdoor-broom-dustpan",
		"https://libmanliquids.com/products/997-wide-commercial-angle-broom",
		"https://libmanliquids.com/products/994-commercial-angle-broom",
		"https://libmanliquids.com/products/1102-fiberforce-outdoor-angle-broom",
		"https://libmanliquids.com/products/499-housekeeper-broom",
		"https://libmanliquids.com/products/1115-wide-commercial-angle-broom-black",
		"https://libmanliquids.com/products/large-precision-angle-broom-clean-fibers-dust-pan",
		"https://libmanliquids.com/products/1086-stiff-sweep-lobby-broom",
		"https://libmanliquids.com/products/502-janitor-corn-broom",
		"https://libmanliquids.com/products/1335-janitor-corn-broom-wood-handle",
		"https://libmanliquids.com/products/919-open-dustpan-with-lobby-broom",
		"https://libmanliquids.com/products/917-closed-dustpan-with-lobby-broom",
		"https://libmanliquids.com/products/lobby-broom-dust-pan-handle-clip-replacement",
		"https://libmanliquids.com/products/1193-deluxe-lobby-dust-pan-broom-closed-lid",
		"https://libmanliquids.com/products/1194-deluxe-open-dustpan-with-lobby-broom",
		"https://libmanliquids.com/products/929-outdoor-scoop",
		"https://libmanliquids.com/products/household-dustpan",
		"https://libmanliquids.com/products/dust-pan-and-brush-set",
		"https://libmanliquids.com/products/526-work-bench-dust-brush",
		"https://libmanliquids.com/products/928-dustpan",
		"https://libmanliquids.com/products/2126-xl-step-on-dustpan",
		"https://libmanliquids.com/products/2125-step-on-dustpan",
		"https://libmanliquids.com/products/907-whisk-broom",
		"https://libmanliquids.com/products/step-on-dustpan",
		"https://libmanliquids.com/products/whisk-broom",
		"https://libmanliquids.com/products/dustpan",
		"https://libmanliquids.com/products/big-dustpan",
		"https://libmanliquids.com/products/911-big-dustpan",
		"https://libmanliquids.com/products/850-heavy-duty-push-broom",
		"https://libmanliquids.com/products/823-multi-surface-heavy-duty-push-broom",
		"https://libmanliquids.com/products/1101-multi-surface-heavy-duty-push-broom",
		"https://libmanliquids.com/products/825-rough-surface-push-broom",
		"https://libmanliquids.com/products/1230-24-fiberforce-multi-surface-push-broom-squeegee",
		"https://libmanliquids.com/products/1293-24-fiberforce-rough-surface-push-broom",
		"https://libmanliquids.com/products/1294-24-fiberforce-smooth-surface-push-broom",
		"https://libmanliquids.com/products/601-60-steel-handle",
		"https://libmanliquids.com/products/602-60-zinc-thread-wood-handle",
		"https://libmanliquids.com/products/1165-60-steel-handle-black",
		"https://libmanliquids.com/products/879-rough-surface-push-broom",
		"https://libmanliquids.com/products/804-multi-surface-push-broom",
		"https://libmanliquids.com/products/878-rough-surface-push-broom",
		"https://libmanliquids.com/products/24-multi-surface-clamp-handle-push-broom",
		"https://libmanliquids.com/products/fiberforce-toilet-brush-caddy",
		"https://libmanliquids.com/products/designer-bowl-brush",
		"https://libmanliquids.com/products/524-all-purpose-scrubbing-brush",
		"https://libmanliquids.com/products/premium-toilet-plunger",
		"https://libmanliquids.com/products/522-long-handle-scrubbing-brush",
		"https://libmanliquids.com/products/510-scrub-brush",
		"https://libmanliquids.com/products/603-48-steel-handle",
		"https://libmanliquids.com/products/513-heavy-duty-scrub-brush",
		"https://libmanliquids.com/products/567-big-scrub-brush",
		"https://libmanliquids.com/products/549-roofing-brush",
		"https://libmanliquids.com/products/toilet-bowl-cleaner",
		"https://libmanliquids.com/products/big-scrub-brush",
		"https://libmanliquids.com/products/short-handle-tampico-scrub-brush",
		"https://libmanliquids.com/products/bathroom-scrubber",
		"https://libmanliquids.com/products/bathroom-scrubber-refills",
		"https://libmanliquids.com/products/floor-scrub-head-only",
		"https://libmanliquids.com/products/scrubster-mop-refill",
		"https://libmanliquids.com/products/3958-gator-mop-with-brush",
		"https://libmanliquids.com/products/gator-mop-with-brush-refill",
		"https://libmanliquids.com/products/988-big-tornado-mop",
		"https://libmanliquids.com/products/big-tornado-mop-refill",
		"https://libmanliquids.com/products/977-cotton-deck-mop",
		"https://libmanliquids.com/products/jumbo-cotton-wet-mop-refill",
		"https://libmanliquids.com/products/90-cotton-deck-mop-refill",
		"https://libmanliquids.com/products/jumbo-cotton-wet-mop",
		"https://libmanliquids.com/products/jumbo-cotton-deck-mop",
		"https://libmanliquids.com/products/944-cotton-deck-mop-refill",
		"https://libmanliquids.com/products/cotton-deck-mop",
		"https://libmanliquids.com/products/982-quick-change-mop-handle",
		"https://libmanliquids.com/products/968-large-blended-looped-end-wet-mop-head-blue",
		"https://libmanliquids.com/products/983-resin-jaw-mop-frame",
		"https://libmanliquids.com/products/972-Large-Cotton-Looped-End-Wet-Mop-Head",
		"https://libmanliquids.com/products/24-Wet-Mop-Head-Cut-End-Cotton",
		"https://libmanliquids.com/products/981-steel-mop-frame-and-handle",
		"https://libmanliquids.com/products/2121-microfiber-looped-end-wet-mop-head-green",
		"https://libmanliquids.com/products/969-Large-Rayon-Looped-End-Wet-Mop-Head",
		"https://libmanliquids.com/products/16-wet-mop-head-cut-end-cotton",
		"https://libmanliquids.com/products/32-wet-mop-head-cut-end-cotton",
		"https://libmanliquids.com/products/mop-bucket-side-press-wringer",
		"https://libmanliquids.com/products/wringer",
		"https://libmanliquids.com/products/1272-utility-bucket",
		"https://libmanliquids.com/products/3-gallon-round-utility-bucket-black",
		"https://libmanliquids.com/products/36-dust-mop",
		"https://libmanliquids.com/products/wet-dry-microfiber-mop-refill",
		"https://libmanliquids.com/products/922-24-dust-mop",
		"https://libmanliquids.com/products/926-extra-large-microfiber-floor-mop",
		"https://libmanliquids.com/products/36-cut-end-dust-mop-head",
		"https://libmanliquids.com/products/24-Cut-End-Dust-Mop-Head",
		"https://libmanliquids.com/products/2-sided-microfiber-mop-refill",
		"https://libmanliquids.com/products/gym-floor-mop",
		"https://libmanliquids.com/products/big-feather-duster",
		"https://libmanliquids.com/products/flexible-microfiber-duster",
		"https://libmanliquids.com/products/585-flexible-microfiber-duster",
		"https://libmanliquids.com/products/590-terry-towels",
		"https://libmanliquids.com/products/all-purpose-reusable-latex-gloves-small",
		"https://libmanliquids.com/products/all-purpose-reusable-latex-gloves-medium",
		"https://libmanliquids.com/products/all-purpose-reusable-latex-gloves-large",
		"https://libmanliquids.com/products/premium-reusable-latex-gloves-small",
		"https://libmanliquids.com/products/premium-reusable-latex-gloves-medium",
		"https://libmanliquids.com/products/premium-reusable-latex-gloves-large",
		"https://libmanliquids.com/products/heavy-duty-reusable-nitrile-gloves-small",
		"https://libmanliquids.com/products/heavy-duty-reusable-nitrile-gloves-medium",
		"https://libmanliquids.com/products/heavy-duty-reusable-nitrile-gloves-large",
		"https://libmanliquids.com/products/591-shop-towels",
		"https://libmanliquids.com/products/microfiber-dusting-mitt",
		"https://libmanliquids.com/products/kitchen-microfiber-cloths",
		"https://libmanliquids.com/products/all-purpose-cleaning-cloth",
		"https://libmanliquids.com/products/1244-industrial-reusable-gloves",
		"https://libmanliquids.com/products/extra-wide-lint-roller",
		"https://libmanliquids.com/products/large-lint-roller-refill",
		"https://libmanliquids.com/products/copper-scrubbers",
		"https://libmanliquids.com/products/528-long-handle-bbq-brush-scraper",
		"https://libmanliquids.com/products/566-extra-long-handle-steel-brush",
		"https://libmanliquids.com/products/529-extra-long-handle-grill-brush-with-scraper",
		"https://libmanliquids.com/products/heavy-duty-scrubbers",
		"https://libmanliquids.com/products/power-scrub-dots-kitchen-bath-sponges-2-pack",
		"https://libmanliquids.com/products/microfiber-sponge-cloths",
		"https://libmanliquids.com/products/568-extra-long-handle-bbq-brush",
		"https://libmanliquids.com/products/heavy-duty-scouring-pads",
		"https://libmanliquids.com/products/stainless-steel-scrubbers",
		"https://libmanliquids.com/products/575-heat-resistant-grill-brush",
		"https://libmanliquids.com/products/595-stainless-steel-grill-brush",
		"https://libmanliquids.com/products/kitchen-vegetable-brush",
		"https://libmanliquids.com/products/dishmatic-dish-wand",
		"https://libmanliquids.com/products/dishmatic-general-purpose-dish-wand-refills-6-pack",
		"https://libmanliquids.com/products/954-extra-wide-floor-squeegee",
		"https://libmanliquids.com/products/515-floor-squeegee",
		"https://libmanliquids.com/products/1276-heavy-duty-squeegee",
		"https://libmanliquids.com/products/542-heavy-duty-curved-floor-squeegee",
		"https://libmanliquids.com/products/24-straight-floor-squeegee-set",
		"https://libmanliquids.com/products/538-straight-floor-squeegee-head",
		"https://libmanliquids.com/products/192-flex-blade-floor-squeegee-head",
		"https://libmanliquids.com/products/539-curved-floor-squeegee-head",
		"https://libmanliquids.com/products/24-flex-blade-floor-squeegee-refill",
		"https://libmanliquids.com/products/window-squeegee",
		"https://libmanliquids.com/products/189-12-stainless-steel-squeegee",
		"https://libmanliquids.com/products/1067-3-in-1-window-squeegee",
		"https://libmanliquids.com/products/190-18-stainless-steel-squeegee",
		"https://libmanliquids.com/products/1061-18-swivel-squeegee",
		"https://libmanliquids.com/products/1060-easy-change-clamp-squeegee",
		"https://libmanliquids.com/products/glass-mirror-cleaner",
		"https://libmanliquids.com/products/all-purpose-cleaner",
		"https://libmanliquids.com/products/600-60-tapered-wood-handle",
		"https://libmanliquids.com/products/607-on-off-flow-thru-handle",
		"https://libmanliquids.com/products/540-vehicle-brush-head",
		"https://libmanliquids.com/products/535-wash-brush-head",
		"https://libmanliquids.com/products/560-vehicle-brush-with-flow-thru-handle",
		"https://libmanliquids.com/products/freedom-concentrated-multi-surface-floor-cleaner-4-pack",
		"https://libmanliquids.com/products/freedom-concentrated-hardwood-cleaner-4-pack",
		"https://libmanliquids.com/products/all-purpose-cleaner-6-pack",
		"https://libmanliquids.com/products/glass-mirror-cleaner-6-pack",
		"https://libmanliquids.com/products/toilet-bowl-cleaner-6-pack",
		"https://libmanliquids.com/products/power-scrub-dots-kitchen-bath-sponge-12-pack",
		"https://libmanliquids.com/products/baked-tough-jobs-sponge-24-pack",
		"https://libmanliquids.com/products/all-purpose-sponge-24-pack",
		"https://libmanliquids.com/products/non-scratch-easy-rinse-sponge-24-pack",
		"https://libmanliquids.com/products/scrub-sponges-suction-hanger-12-pack",
		"https://libmanliquids.com/products/heavy-duty-wonder-mop",
		"https://libmanliquids.com/products/heavy-duty-wonder-mop-refill",
		"https://libmanliquids.com/products/maid-caddy",
		"https://libmanliquids.com/products/516-dual-surface-scrub-brush-head",
		"https://libmanliquids.com/products/955-roller-mop-scrub-brush",
		"https://libmanliquids.com/products/4-gallon-clean-rinse-bucket-2-pack",
		"https://libmanliquids.com/products/tornado-spin-mop-system",
		"https://libmanliquids.com/products/tornado-spin-mop-system-refill",
		"https://libmanliquids.com/products/1572-32-gallon-trash-can",
		"https://libmanliquids.com/products/1575-32-gallon-trash-can-lid-green",
		"https://libmanliquids.com/products/32-gallon-trash-can-lid-black",
		"https://libmanliquids.com/products/1464-32-gallon-trash-can-lid-grey",
		"https://libmanliquids.com/products/32-gallon-trash-can-lid-green",
		"https://libmanliquids.com/products/1573-32-gallon-trash-can-lid-gray",
		"https://libmanliquids.com/products/1571-32-gallon-trash-can-lid-black",
		"https://libmanliquids.com/products/1574-32-gallon-trash-can-green",
		"https://libmanliquids.com/products/1570-32-gallon-trash-can-black",
		"https://libmanliquids.com/products/1262-microfiber-cleaning-cloths",
		"https://libmanliquids.com/products/580-microfiber-cleaning-cloths",
		"https://libmanliquids.com/products/1576-industrial-heavy-duty-floor-scrub",
		"https://libmanliquids.com/products/clean-fibers-dustpan",
		"https://libmanliquids.com/products/non-scratch-easy-rinse-sponges-9-pack",
		"https://libmanliquids.com/products/heavy-duty-easy-rinse-sponge-9-pack",
		"https://libmanliquids.com/products/rinse-n-wring-mop-system",
		"https://libmanliquids.com/products/1503-24-contractor-grade-multi-surface-push-broom-fiberglass-handle",
		"https://libmanliquids.com/products/1559-swivel-grout-scrub-brush",
		"https://libmanliquids.com/products/1616-swivel-grout-scrub-brush-head-only",
		"https://libmanliquids.com/products/1683-60-threaded-steel-handle-no-hex",
		"https://libmanliquids.com/products/1681-fiberforce-all-purpose-floor-scrub",
		"https://libmanliquids.com/products/618-52-taper-threaded-handle",
		"https://libmanliquids.com/products/rinse-n-wring-microfiber-mop-system-refill",
		"https://libmanliquids.com/products/petplus-angle-broom-dustpan",
		"https://libmanliquids.com/products/iluma-glass-mirror-concentrated-cleaning-system",
		"https://libmanliquids.com/products/iluma-glass-mirror-concentrated-cleaning-refills",
		"https://libmanliquids.com/products/7552-24-soft-rubber-floor-replacement-squeegee-head",
		"https://libmanliquids.com/products/freedom-dual-sided-microfiber-spray-mop",
		"https://libmanliquids.com/products/dual-sided-freedom-spray-mop-refill",
		"https://libmanliquids.com/products/tear-n-wipe-cloths",
		"https://libmanliquids.com/products/all-purpose-spray-bottle-0",
		"https://libmanliquids.com/products/non-scratch-scouring-pads",
		"https://libmanliquids.com/products/1811-two-sided-caution-wet-floor-sign-clip",
		"https://libmanliquids.com/products/1243-small-blended-looped-end-wet-mop-head-blue",
		"https://libmanliquids.com/products/9002280-pyramid-display",
		"https://libmanliquids.com/products/1810-12-rough-surface-angle-broom",
		"https://libmanliquids.com/products/leather-conditioner-45oz",
		"https://libmanliquids.com/products/leather-conditioner-16oz",
		"https://libmanliquids.com/products/oakwood-leather-oil",
		"https://libmanliquids.com/products/oakwood-glycerine-leather-cleaner",
		"https://libmanliquids.com/products/oakwood-liquid-saddle-soap",
		"https://libmanliquids.com/products/pin-and-bristle-brush",
		"https://libmanliquids.com/products/loose-hair-remover-glove",
		"https://libmanliquids.com/products/detangling-slicker-brush",
		"https://libmanliquids.com/products/step-n-stand-dustpan",
		"https://libmanliquids.com/products/large-precision-angle-broom-step-n-stand-dustpan",
		"https://libmanliquids.com/products/lightning-spin-mop-system",
		"https://libmanliquids.com/products/bucket-trolley",
		"https://libmanliquids.com/products/1786-16-wet-mop-head-cut-end-cotton-brick-pack",
		"https://libmanliquids.com/products/1787-20-wet-mop-head-cut-end-cotton-brick-pack",
		"https://libmanliquids.com/products/1788-24-wet-mop-head-cut-end-cotton-brick-pack",
		"https://libmanliquids.com/products/1828-24-wet-mop-head-scrub-pad-cut-end-cotton-brick-pack",
		"https://libmanliquids.com/products/1789-32-wet-mop-head-cut-end-cotton-brick-pack",
		"https://libmanliquids.com/products/1790-small-cotton-looped-end-wet-mop-head-brick-pack",
		"https://libmanliquids.com/products/1791-medium-cotton-looped-end-wet-mop-head-brick-pack",
		"https://libmanliquids.com/products/1792-large-cotton-looped-end-wet-mop-head-brick-pack",
		"https://libmanliquids.com/products/1830-large-cotton-looped-end-wet-mop-head-scrub-pad-brick-pack",
		"https://libmanliquids.com/products/1793-x-large-cotton-looped-end-wet-mop-head-brick-pack",
		"https://libmanliquids.com/products/1794-small-blended-looped-end-wet-mop-head-blue-brick-pack",
		"https://libmanliquids.com/products/1795-medium-blended-looped-end-wet-mop-head-blue-brick-pack",
		"https://libmanliquids.com/products/1796-large-blended-looped-end-wet-mop-head-blue-brick-pack",
		"https://libmanliquids.com/products/1829-large-blended-looped-end-wet-mop-head-scrub-pad-blue-brick-pack",
		"https://libmanliquids.com/products/1797-x-large-blended-looped-end-wet-mop-head-blue-brick-pack",
		"https://libmanliquids.com/products/1802-small-rayon-looped-end-wet-mop-head-bluewhite-brick-pack",
		"https://libmanliquids.com/products/1803-medium-rayon-looped-end-wet-mop-head-bluewhite-brick-pack",
		"https://libmanliquids.com/products/1804-large-rayon-looped-end-wet-mop-head-bluewhite-brick-pack",
		"https://libmanliquids.com/products/1805-x-large-rayon-looped-end-wet-mop-head-bluewhite-brick-pack",
		"https://libmanliquids.com/products/1798-small-premium-green-blend-looped-end-wet-mop-head-green-brick-pack",
		"https://libmanliquids.com/products/1799-medium-premium-green-blend-looped-end-wet-mop-head-green-brick-pack",
		"https://libmanliquids.com/products/1800-large-premium-green-blend-looped-end-wet-mop-head-green-brick-pack",
		"https://libmanliquids.com/products/1801-x-large-premium-green-blend-looped-end-wet-mop-head-green-brick-pack",
		"https://libmanliquids.com/products/1831-large-microfiber-looped-end-wet-mop-head-scrub-pad-green",
		"https://libmanliquids.com/products/1832-medium-microfiber-looped-end-wet-mop-head-blue",
		"https://libmanliquids.com/products/1833-large-microfiber-looped-end-wet-mop-head-blue",
		"https://libmanliquids.com/products/1855-microfiber-cleaning-cloths",
		"https://libmanliquids.com/products/triplegrip-microfiber-scrub-cloths",
		"https://libmanliquids.com/products/1860-glass-mirror-lint-free-cloths",
		"https://libmanliquids.com/products/lightning-spin-mop-system-refill",
		"https://libmanliquids.com/products/pro-grade-microfiber-spin-mop-system",
		"https://libmanliquids.com/products/pro-grade-microfiber-spin-mop-system-refill",
		"https://libmanliquids.com/products/1846-48-cotton-blend-dust-mop-head-48oz",
		"https://libmanliquids.com/products/1847-24-microfiber-dust-mop-head-8oz",
		"https://libmanliquids.com/products/1848-36-microfiber-dust-mop-head-10oz",
		"https://libmanliquids.com/products/1849-48-microfiber-dust-mop-head-13oz",
	} // URL to fetch HTML content from
	localFilePath := "libmanliquids.html" // Path where HTML file will be stored

	var getData []string

	for _, urls := range remoteAPIURL {
		getData = append(getData, getDataFromURL(urls)) // If not, download HTML content from URL
	}
	appendAndWriteToFile(localFilePath, strings.Join(getData, "")) // Save downloaded content to file

	finalList := extractPDFUrls(strings.Join(getData, "")) // Extract all PDF links from HTML content

	outputDir := "PDFs/" // Directory to store downloaded PDFs

	if !directoryExists(outputDir) { // Check if directory exists
		createDirectory(outputDir, 0o755) // Create directory with read-write-execute permissions
	}

	// Remove duplicates from a given slice.
	finalList = removeDuplicatesFromSlice(finalList)

	// Loop through all extracted PDF URLs
	for _, urls := range finalList {
		if !hasDomain(urls) {
			urls = "https://libmanliquids.com" + urls

		}
		if isUrlValid(urls) { // Check if the final URL is valid
			downloadPDF(urls, outputDir) // Download the PDF
		}
	}
}

// hasDomain checks if the given string has a domain (host part)
func hasDomain(rawURL string) bool {
	// Try parsing the raw string as a URL
	parsed, err := url.Parse(rawURL)
	if err != nil { // If parsing fails, it's not a valid URL
		return false
	}
	// If the parsed URL has a non-empty Host, then it has a domain/host
	return parsed.Host != ""
}

// Opens a file in append mode, or creates it, and writes the content to it
func appendAndWriteToFile(path string, content string) {
	filePath, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644) // Open file with specified flags and permissions
	if err != nil {
		log.Println(err) // Log error if opening fails
	}
	_, err = filePath.WriteString(content + "\n") // Write content to file
	if err != nil {
		log.Println(err) // Log error if writing fails
	}
	err = filePath.Close() // Close the file
	if err != nil {
		log.Println(err) // Log error if closing fails
	}
}

// Extracts filename from full path (e.g. "/dir/file.pdf" → "file.pdf")
func getFilename(path string) string {
	return filepath.Base(path) // Use Base function to get file name only
}

// Converts a raw URL into a sanitized PDF filename safe for filesystem
func urlToFilename(rawURL string) string {
	lower := strings.ToLower(rawURL) // Convert URL to lowercase
	lower = getFilename(lower)       // Extract filename from URL

	reNonAlnum := regexp.MustCompile(`[^a-z0-9]`)   // Regex to match non-alphanumeric characters
	safe := reNonAlnum.ReplaceAllString(lower, "_") // Replace non-alphanumeric with underscores

	safe = regexp.MustCompile(`_+`).ReplaceAllString(safe, "_") // Collapse multiple underscores into one
	safe = strings.Trim(safe, "_")                              // Trim leading and trailing underscores

	var invalidSubstrings = []string{
		"_pdf", // Substring to remove from filename
	}

	for _, invalidPre := range invalidSubstrings { // Remove unwanted substrings
		safe = removeSubstring(safe, invalidPre)
	}

	if getFileExtension(safe) != ".pdf" { // Ensure file ends with .pdf
		safe = safe + ".pdf"
	}

	return safe // Return sanitized filename
}

// Removes all instances of a specific substring from input string
func removeSubstring(input string, toRemove string) string {
	result := strings.ReplaceAll(input, toRemove, "") // Replace substring with empty string
	return result
}

// Gets the file extension from a given file path
func getFileExtension(path string) string {
	return filepath.Ext(path) // Extract and return file extension
}

// Checks if a file exists at the specified path
func fileExists(filename string) bool {
	info, err := os.Stat(filename) // Get file info
	if err != nil {                // If error occurs, file doesn't exist
		return false
	}
	return !info.IsDir() // Return true if path is a file (not a directory)
}

// Downloads a PDF from given URL and saves it in the specified directory
func downloadPDF(finalURL, outputDir string) bool {
	filename := strings.ToLower(urlToFilename(finalURL)) // Sanitize the filename
	filePath := filepath.Join(outputDir, filename)       // Construct full path for output file

	if fileExists(filePath) { // Skip if file already exists
		log.Printf("File already exists, skipping: %s", filePath)
		return false
	}

	client := &http.Client{Timeout: 15 * time.Minute} // Create HTTP client with timeout

	resp, err := client.Get(finalURL) // Send HTTP GET request
	if err != nil {
		log.Printf("Failed to download %s: %v", finalURL, err)
		return false
	}
	defer resp.Body.Close() // Ensure response body is closed

	if resp.StatusCode != http.StatusOK { // Check if response is 200 OK
		log.Printf("Download failed for %s: %s", finalURL, resp.Status)
		return false
	}

	contentType := resp.Header.Get("Content-Type")                                                                  // Get content type of response
	if !strings.Contains(contentType, "binary/octet-stream") && !strings.Contains(contentType, "application/pdf") { // Check if it's a PDF
		log.Printf("Invalid content type for %s: %s (expected binary/octet-stream) (expected application/pdf)", finalURL, contentType)
		return false
	}

	var buf bytes.Buffer                     // Create a buffer to hold response data
	written, err := io.Copy(&buf, resp.Body) // Copy data into buffer
	if err != nil {
		log.Printf("Failed to read PDF data from %s: %v", finalURL, err)
		return false
	}
	if written == 0 { // Skip empty files
		log.Printf("Downloaded 0 bytes for %s; not creating file", finalURL)
		return false
	}

	out, err := os.Create(filePath) // Create output file
	if err != nil {
		log.Printf("Failed to create file for %s: %v", finalURL, err)
		return false
	}
	defer out.Close() // Ensure file is closed after writing

	if _, err := buf.WriteTo(out); err != nil { // Write buffer contents to file
		log.Printf("Failed to write PDF to file for %s: %v", finalURL, err)
		return false
	}

	log.Printf("Successfully downloaded %d bytes: %s → %s", written, finalURL, filePath) // Log success
	return true
}

// Checks whether a given directory exists
func directoryExists(path string) bool {
	directory, err := os.Stat(path) // Get info for the path
	if err != nil {
		return false // Return false if error occurs
	}
	return directory.IsDir() // Return true if it's a directory
}

// Creates a directory at given path with provided permissions
func createDirectory(path string, permission os.FileMode) {
	err := os.Mkdir(path, permission) // Attempt to create directory
	if err != nil {
		log.Println(err) // Log error if creation fails
	}
}

// Verifies whether a string is a valid URL format
func isUrlValid(uri string) bool {
	_, err := url.ParseRequestURI(uri) // Try parsing the URL
	return err == nil                  // Return true if valid
}

// Removes duplicate strings from a slice
func removeDuplicatesFromSlice(slice []string) []string {
	check := make(map[string]bool) // Map to track seen values
	var newReturnSlice []string    // Slice to store unique values
	for _, content := range slice {
		if !check[content] { // If not already seen
			check[content] = true                            // Mark as seen
			newReturnSlice = append(newReturnSlice, content) // Add to result
		}
	}
	return newReturnSlice
}

// Extracts all links to PDF files from given HTML string
func extractPDFUrls(htmlInput string) []string {
	var pdfLinks []string // Slice to hold found PDF links

	doc, err := html.Parse(strings.NewReader(htmlInput)) // Parse HTML content
	if err != nil {
		log.Println(err) // Log parse error
		return nil
	}

	var traverse func(*html.Node) // Recursive function to traverse HTML nodes
	traverse = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" { // If it's an <a> tag
			for _, attr := range n.Attr {
				if attr.Key == "href" { // Look for href attribute
					href := strings.TrimSpace(attr.Val)                  // Get link
					if strings.Contains(strings.ToLower(href), ".pdf") { // If link points to a PDF
						pdfLinks = append(pdfLinks, href) // Add to list
					}
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling { // Traverse children
			traverse(c)
		}
	}

	traverse(doc)   // Start traversal from root
	return pdfLinks // Return found PDF links
}

// Performs HTTP GET request and returns response body as string
func getDataFromURL(uri string) string {
	log.Println("Scraping", uri)   // Log which URL is being scraped
	response, err := http.Get(uri) // Send GET request
	if err != nil {
		log.Println(err) // Log if request fails
	}

	body, err := io.ReadAll(response.Body) // Read the body of the response
	if err != nil {
		log.Println(err) // Log read error
	}

	err = response.Body.Close() // Close response body
	if err != nil {
		log.Println(err) // Log error during close
	}
	return string(body) // Return response body as string
}
