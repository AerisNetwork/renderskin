package renderskin

import "image"

var (
	// HeadTop is the top side of the head
	HeadTop = image.Rect(8, 0, 16, 8)
	// HeadBottom is the bottom side of the head
	HeadBottom = image.Rect(16, 0, 24, 8)
	// HeadRight is the right side of the head
	HeadRight = image.Rect(0, 8, 8, 16)
	// HeadFront is the front side of the head
	HeadFront = image.Rect(8, 8, 16, 16)
	// HeadLeft is the left side of the head
	HeadLeft = image.Rect(16, 8, 24, 16)
	// HeadBack is the back side of the head
	HeadBack = image.Rect(24, 8, 32, 16)
	// HeadOverlayTop is the top side of the head overlay
	HeadOverlayTop = image.Rect(40, 0, 48, 8)
	// HeadOverlayBottom is the bottom side of the head overlay
	HeadOverlayBottom = image.Rect(48, 0, 56, 8)
	// HeadOverlayRight is the right side of the head overlay
	HeadOverlayRight = image.Rect(32, 8, 40, 16)
	// HeadOverlayFront is the front side of the head overlay
	HeadOverlayFront = image.Rect(40, 8, 48, 16)
	// HeadOverlayLeft is the left side of the head overlay
	HeadOverlayLeft = image.Rect(48, 8, 56, 16)
	// HeadOverlayBack is the back side of the head overlay
	HeadOverlayBack = image.Rect(56, 8, 64, 16)
	// RightLegTop is the top side of the right leg
	RightLegTop = image.Rect(4, 16, 8, 20)
	// RightLegBottom is the bottom side of the right leg
	RightLegBottom = image.Rect(8, 16, 12, 20)
	// RightLegRight is the right side of the right leg
	RightLegRight = image.Rect(0, 20, 4, 32)
	// RightLegFront is the front side of the right leg
	RightLegFront = image.Rect(4, 20, 8, 32)
	// RightLegLeft is the left side of the right leg
	RightLegLeft = image.Rect(8, 20, 12, 32)
	// RightLegBack is the back side of the right leg
	RightLegBack = image.Rect(12, 20, 16, 32)
	// TorsoTop is the top side of the torso
	TorsoTop = image.Rect(20, 16, 28, 20)
	// TorsoBottom is the bottom side of the torso
	TorsoBottom = image.Rect(28, 16, 36, 20)
	// TorsoRight is the right side of the torso
	TorsoRight = image.Rect(16, 20, 20, 32)
	// TorsoFront is the front side of the torso
	TorsoFront = image.Rect(20, 20, 28, 32)
	// TorsoLeft is the left side of the torso
	TorsoLeft = image.Rect(28, 20, 32, 32)
	// TorsoBack is the back side of the torso
	TorsoBack = image.Rect(32, 20, 40, 32)
	// RightArmTopRegular is the top side of the right arm for regular skin models
	RightArmTopRegular = image.Rect(44, 16, 48, 20)
	// RightArmTopSlim is the top side of the right arm for slim skin models
	RightArmTopSlim = image.Rect(44, 16, 47, 20)
	// RightArmBottomRegular is the bottom side of the right arm for regular skin models
	RightArmBottomRegular = image.Rect(48, 16, 52, 20)
	// RightArmBottomSlim is the bottom side of the right arm for slim skin models
	RightArmBottomSlim = image.Rect(47, 16, 50, 20)
	// RightArmRight is the right side of the right arm
	RightArmRight = image.Rect(40, 20, 44, 32)
	// RightArmFrontRegular is the front side of the right arm for regular skin models
	RightArmFrontRegular = image.Rect(44, 20, 48, 32)
	// RightArmFrontSlim is the front side of the right arm for slim skin models
	RightArmFrontSlim = image.Rect(44, 20, 47, 32)
	// RightArmLeftRegular is the left side of the right arm for regular skin models
	RightArmLeftRegular = image.Rect(48, 20, 52, 32)
	// RightArmLeftSlim is the left side of the right arm for slim skin models
	RightArmLeftSlim = image.Rect(47, 20, 51, 32)
	// RightArmBackRegular is the back side of the right arm for regular skin models
	RightArmBackRegular = image.Rect(52, 20, 56, 32)
	// RightArmBackSlim is the back side of the right arm for slim skin models
	RightArmBackSlim = image.Rect(51, 20, 54, 32)
	// LeftLegTop is the top side of the left leg
	LeftLegTop = image.Rect(20, 48, 24, 52)
	// LeftLegBottom is the bottom side of the left leg
	LeftLegBottom = image.Rect(24, 48, 28, 52)
	// LeftLegRight is the right side of the left leg
	LeftLegRight = image.Rect(16, 52, 20, 64)
	// LeftLegFront is the front side of the left leg
	LeftLegFront = image.Rect(20, 52, 24, 64)
	// LeftLegLeft is the left side of the left leg
	LeftLegLeft = image.Rect(24, 52, 28, 64)
	// LeftLegBack is the back side of the left leg
	LeftLegBack = image.Rect(28, 52, 32, 64)
	// LeftArmTopRegular is the top side of the left arm for regular skin models
	LeftArmTopRegular = image.Rect(36, 48, 40, 52)
	// LeftArmTopSlim is the top side of the left arm for slim skin models
	LeftArmTopSlim = image.Rect(36, 48, 39, 52)
	// LeftArmBottomRegular is the bottom side of the left arm for regular skin models
	LeftArmBottomRegular = image.Rect(40, 48, 44, 52)
	// LeftArmBottomSlim is the bottom side of the left arm for slim skin models
	LeftArmBottomSlim = image.Rect(39, 48, 42, 52)
	// LeftArmRight is the right side of the left arm
	LeftArmRight = image.Rect(32, 52, 36, 64)
	// LeftArmFrontRegular is the front side of the left arm for regular skin models
	LeftArmFrontRegular = image.Rect(36, 52, 40, 64)
	// LeftArmFrontSlim is the front side of the left arm for slim skin models
	LeftArmFrontSlim = image.Rect(36, 52, 39, 64)
	// LeftArmLeftRegular is the left side of the left arm for regular skin models
	LeftArmLeftRegular = image.Rect(40, 52, 44, 64)
	// LeftArmLeftSlim is the left side of the left arm for slim skin models
	LeftArmLeftSlim = image.Rect(39, 52, 43, 64)
	// LeftArmBackRegular is the back side of the left arm for regular skin models
	LeftArmBackRegular = image.Rect(44, 52, 48, 64)
	// LeftArmBackSlim is the back side of the left arm for slim skin models
	LeftArmBackSlim = image.Rect(43, 52, 46, 64)
	// RightLegOverlayTop is the top side of the right leg overlay
	RightLegOverlayTop = image.Rect(4, 48, 8, 52)
	// RightLegOverlayBottom is the bottom side of the right leg overlay
	RightLegOverlayBottom = image.Rect(8, 48, 12, 52)
	// RightLegOverlayRight is the right side of the right leg overlay
	RightLegOverlayRight = image.Rect(0, 36, 4, 48)
	// RightLegOverlayFront is the front side of the right leg overlay
	RightLegOverlayFront = image.Rect(4, 36, 8, 48)
	// RightLegOverlayLeft is the left side of the right leg overlay
	RightLegOverlayLeft = image.Rect(8, 36, 12, 48)
	// RightLegOverlayBack is the back side of the right leg overlay
	RightLegOverlayBack = image.Rect(12, 36, 16, 48)
	// TorsoOverlayTop is the top side of the torso overlay
	TorsoOverlayTop = image.Rect(20, 48, 28, 52)
	// TorsoOverlayBottom is the bottom side of the torso overlay
	TorsoOverlayBottom = image.Rect(28, 48, 36, 52)
	// TorsoOverlayRight is the right side of the torso overlay
	TorsoOverlayRight = image.Rect(16, 36, 20, 48)
	// TorsoOverlayFront is the front side of the torso overlay
	TorsoOverlayFront = image.Rect(20, 36, 28, 48)
	// TorsoOverlayLeft is the left side of the torso overlay
	TorsoOverlayLeft = image.Rect(28, 36, 32, 48)
	// TorsoOverlayBack is the back side of the torso overlay
	TorsoOverlayBack = image.Rect(32, 36, 40, 48)
	// RightArmOverlayTopRegular is the top side of the right arm overlay for regular skin models
	RightArmOverlayTopRegular = image.Rect(44, 48, 48, 52)
	// RightArmOverlayTopSlim is the top side of the right arm overlay for slim skin models
	RightArmOverlayTopSlim = image.Rect(44, 48, 47, 52)
	// RightArmOverlayBottomRegular is the bottom side of the right arm overlay for regular skin models
	RightArmOverlayBottomRegular = image.Rect(48, 48, 52, 52)
	// RightArmOverlayBottomSlim is the bottom side of the right arm overlay for slim skin models
	RightArmOverlayBottomSlim = image.Rect(47, 48, 50, 52)
	// RightArmOverlayRight is the right side of the right arm overlay
	RightArmOverlayRight = image.Rect(40, 36, 44, 48)
	// RightArmOverlayFrontRegular is the front side of the right arm overlay for regular skin models
	RightArmOverlayFrontRegular = image.Rect(44, 36, 48, 48)
	// RightArmOverlayFrontSlim is the front side of the right arm overlay for slim skin models
	RightArmOverlayFrontSlim = image.Rect(44, 36, 47, 48)
	// RightArmOverlayLeftRegular is the left side of the right arm overlay for regular skin models
	RightArmOverlayLeftRegular = image.Rect(48, 36, 52, 48)
	// RightArmOverlayLeftSlim is the left side of the right arm overlay for slim skin models
	RightArmOverlayLeftSlim = image.Rect(47, 36, 51, 48)
	// RightArmOverlayBackRegular is the back side of the right arm overlay for regular skin models
	RightArmOverlayBackRegular = image.Rect(52, 36, 56, 48)
	// RightArmOverlayBackSlim is the back side of the right arm overlay for slim skin models
	RightArmOverlayBackSlim = image.Rect(51, 36, 54, 48)
	// LeftLegOverlayTop is the top side of the left leg overlay
	LeftLegOverlayTop = image.Rect(4, 48, 8, 52)
	// LeftLegOverlayBottom is the bottom side of the left leg overlay
	LeftLegOverlayBottom = image.Rect(8, 48, 12, 52)
	// LeftLegOverlayRight is the right side of the left leg overlay
	LeftLegOverlayRight = image.Rect(0, 52, 4, 64)
	// LeftLegOverlayFront is the front side of the left leg overlay
	LeftLegOverlayFront = image.Rect(4, 52, 8, 64)
	// LeftLegOverlayLeft is the left side of the left leg overlay
	LeftLegOverlayLeft = image.Rect(8, 52, 12, 64)
	// LeftLegOverlayBack is the back side of the left leg overlay
	LeftLegOverlayBack = image.Rect(12, 52, 16, 64)
	// LeftArmOverlayTopRegular is the top side of the left arm overlay for regular skin models
	LeftArmOverlayTopRegular = image.Rect(52, 48, 56, 52)
	// LeftArmOverlayTopSlim is the top side of the left arm overlay for slim skin models
	LeftArmOverlayTopSlim = image.Rect(52, 48, 55, 52)
	// LeftArmOverlayBottomRegular is the bottom side of the left arm overlay for regular skin models
	LeftArmOverlayBottomRegular = image.Rect(56, 48, 60, 52)
	// LeftArmOverlayBottomSlim is the bottom side of the left arm overlay for slim skin models
	LeftArmOverlayBottomSlim = image.Rect(55, 48, 58, 52)
	// LeftArmOverlayRight is the right side of the left arm overlay
	LeftArmOverlayRight = image.Rect(48, 52, 52, 64)
	// LeftArmOverlayFrontRegular is the front side of the left arm overlay for regular skin models
	LeftArmOverlayFrontRegular = image.Rect(52, 52, 56, 64)
	// LeftArmOverlayFrontSlim is the front side of the left arm overlay for slim skin models
	LeftArmOverlayFrontSlim = image.Rect(52, 52, 55, 64)
	// LeftArmOverlayLeftRegular is the left side of the left arm overlay for regular skin models
	LeftArmOverlayLeftRegular = image.Rect(56, 52, 60, 64)
	// LeftArmOverlayLeftSlim is the left side of the left arm overlay for slim skin models
	LeftArmOverlayLeftSlim = image.Rect(55, 52, 59, 64)
	// LeftArmOverlayBackRegular is the back side of the left arm overlay for regular skin models
	LeftArmOverlayBackRegular = image.Rect(60, 52, 64, 64)
	// LeftArmOverlayBackSlim is the back side of the left arm overlay for slim skin models
	LeftArmOverlayBackSlim = image.Rect(59, 52, 62, 64)
)

// rightArmTop returns the top of a right arm based on if the skin is slim or not
func rightArmTop(slim bool) image.Rectangle {
	if slim {
		return RightArmTopSlim
	}
	return RightArmTopRegular
}

// rightArmBottom returns the bottom of a right arm based on if the skin is slim or not
func rightArmBottom(slim bool) image.Rectangle {
	if slim {
		return RightArmBottomSlim
	}
	return RightArmBottomRegular
}

// rightArmFront returns the front of a right arm based on if the skin is slim or not
func rightArmFront(slim bool) image.Rectangle {
	if slim {
		return RightArmFrontSlim
	}
	return RightArmFrontRegular
}

// rightArmLeft returns the left of a right arm based on if the skin is slim or not
func rightArmLeft(slim bool) image.Rectangle {
	if slim {
		return RightArmLeftSlim
	}
	return RightArmLeftRegular
}

// rightArmBack returns the back of a right arm based on if the skin is slim or not
func rightArmBack(slim bool) image.Rectangle {
	if slim {
		return RightArmBackSlim
	}
	return RightArmBackRegular
}

// leftArmTop returns the top of a left arm based on if the skin is slim or not
func leftArmTop(slim bool) image.Rectangle {
	if slim {
		return LeftArmTopSlim
	}
	return LeftArmTopRegular
}

// leftArmBottom returns the bottom of a left arm based on if the skin is slim or not
func leftArmBottom(slim bool) image.Rectangle {
	if slim {
		return LeftArmBottomSlim
	}
	return LeftArmBottomRegular
}

// leftArmFront returns the front of a left arm based on if the skin is slim or not
func leftArmFront(slim bool) image.Rectangle {
	if slim {
		return LeftArmFrontSlim
	}
	return LeftArmFrontRegular
}

// leftArmLeft returns the left of a left arm based on if the skin is slim or not
func leftArmLeft(slim bool) image.Rectangle {
	if slim {
		return LeftArmLeftSlim
	}
	return LeftArmLeftRegular
}

// leftArmBack returns the back of a left arm based on if the skin is slim or not
func leftArmBack(slim bool) image.Rectangle {
	if slim {
		return LeftArmBackSlim
	}
	return LeftArmBackRegular
}

// rightArmOverlayTop returns the top of a right arm overlay based on if the skin is slim or not
func rightArmOverlayTop(slim bool) image.Rectangle {
	if slim {
		return RightArmOverlayTopSlim
	}
	return RightArmOverlayTopRegular
}

// rightArmOverlayBottom returns the bottom of a right arm overlay based on if the skin is slim or not
func rightArmOverlayBottom(slim bool) image.Rectangle {
	if slim {
		return RightArmOverlayBottomSlim
	}
	return RightArmOverlayBottomRegular
}

// rightArmOverlayFront returns the front of a right arm overlay based on if the skin is slim or not
func rightArmOverlayFront(slim bool) image.Rectangle {
	if slim {
		return RightArmOverlayFrontSlim
	}
	return RightArmOverlayFrontRegular
}

// rightArmOverlayLeft returns the left of a right arm overlay based on if the skin is slim or not
func rightArmOverlayLeft(slim bool) image.Rectangle {
	if slim {
		return RightArmOverlayLeftSlim
	}
	return RightArmOverlayLeftRegular
}

// rightArmOverlayBack returns the back of a right arm overlay based on if the skin is slim or not
func rightArmOverlayBack(slim bool) image.Rectangle {
	if slim {
		return RightArmOverlayBackSlim
	}
	return RightArmOverlayBackRegular
}

// leftArmOverlayTop returns the top of a left arm overlay based on if the skin is slim or not
func leftArmOverlayTop(slim bool) image.Rectangle {
	if slim {
		return LeftArmOverlayTopSlim
	}
	return LeftArmOverlayTopRegular
}

// leftArmOverlayBottom returns the bottom of a left arm overlay based on if the skin is slim or not
func leftArmOverlayBottom(slim bool) image.Rectangle {
	if slim {
		return LeftArmOverlayBottomSlim
	}
	return LeftArmOverlayBottomRegular
}

// leftArmOverlayFront returns the front of a left arm overlay based on if the skin is slim or not
func leftArmOverlayFront(slim bool) image.Rectangle {
	if slim {
		return LeftArmOverlayFrontSlim
	}
	return LeftArmOverlayFrontRegular
}

// leftArmOverlayLeft returns the left of a left arm overlay based on if the skin is slim or not
func leftArmOverlayLeft(slim bool) image.Rectangle {
	if slim {
		return LeftArmOverlayLeftSlim
	}
	return LeftArmOverlayLeftRegular
}

// leftArmOverlayBack returns the back of a left arm overlay based on if the skin is slim or not
func leftArmOverlayBack(slim bool) image.Rectangle {
	if slim {
		return LeftArmOverlayBackSlim
	}
	return LeftArmOverlayBackRegular
}
