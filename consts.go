package main

var AnimationsFolder = Folder{
	"Animations",
	[]Folder{
		{
			"AnimationClip",
			nil,
		},
		{
			"Animator",
			nil,
		},
	},
}

var AudioFolder = Folder{
	"Audio",
	[]Folder{
		{
			"AudioClip",
			nil,
		},
		{
			"AudioMixers",
			nil,
		},
	},
}

var MaterialsFolder = Folder{
	"Materials",
	[]Folder{
		{
			"BaseMaterial",
			nil,
		},
	},
}

var ShadersFolder = Folder{
	"Shaders",
	[]Folder{
		{
			"BaseShader",
			nil,
		},
	},
}

var ArtSectionFolder = Folder{
	"Art",
	[]Folder{
		AnimationsFolder,
		AudioFolder,
		MaterialsFolder,
		ShadersFolder,
		{
			"Models",
			nil,
		},
		{
			"PhysicsMaterials",
			nil,
		},
	},
}

var BaseModulesFolder = []Folder{
	ArtSectionFolder,
	{
		"Script",
		nil,
	},
	{
		"Prefab",
		nil,
	},
}

var InitialModulesFolder = Folder{
	name: "_Modules",
	children: []Folder{
		{
			"_Los&Found",
			nil,
		},
		{
			"Player",
			BaseModulesFolder,
		},
		{
			"HUD",
			BaseModulesFolder,
		},
		{
			"Environment",
			BaseModulesFolder,
		},
		{
			"ItemsDrop",
			nil,
		},
	},
}

var InitialToolsFolder = Folder{
	name: "Tools",
	children: []Folder{
		{
			"Editor",
			nil,
		},
	},
}

var initialFileStruct = []Folder{
	InitialModulesFolder,
	{name: "Documentation"},
	{name: "Plugins"},
	InitialToolsFolder,
	{name: "Settings"},
}
