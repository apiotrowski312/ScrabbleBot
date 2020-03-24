package gaddag

import (
	"reflect"
	"testing"
)

func Test_gaddag_add_one_word(t *testing.T) {
	gd := node{
		children: map[rune]node{},
	}

	wNode := node{
		isWord: false,
		children: map[rune]node{
			'.': node{
				isWord: false,
				children: map[rune]node{
					'o': node{
						isWord: false,
						children: map[rune]node{
							'r': node{
								isWord: false,
								children: map[rune]node{
									'd': node{
										isWord:   true,
										children: map[rune]node{},
									},
								},
							},
						},
					},
				},
			},
		},
	}

	rNode := node{
		isWord: false,
		children: map[rune]node{
			'o': node{
				isWord: false,
				children: map[rune]node{
					'w': node{
						isWord: false,
						children: map[rune]node{
							'.': node{
								isWord: false,
								children: map[rune]node{
									'd': node{
										isWord:   true,
										children: map[rune]node{},
									},
								},
							},
						},
					},
				},
			},
		},
	}

	t.Run("Test adding one word", func(t *testing.T) {
		gd.addWord("word")

		if reflect.DeepEqual(gd.children['w'], wNode) != true {
			t.Errorf("Adding new word to graph has errors. Node: %v is different from example: %v", gd.children['w'], wNode)
			return
		}

		if reflect.DeepEqual(gd.children['r'], rNode) != true {
			t.Errorf("Adding new word to graph has errors. Node: %v is different from example: %v", gd.children['r'], rNode)
			return
		}

	})
}

func Test_gaddag_add_five_words(t *testing.T) {
	gd := node{
		children: map[rune]node{},
	}

	wNode := node{
		isWord: false,
		children: map[rune]node{
			'.': node{
				isWord: false,
				children: map[rune]node{
					'o': node{
						isWord: false,
						children: map[rune]node{
							'r': node{
								isWord: false,
								children: map[rune]node{
									'k': node{
										isWord:   true,
										children: map[rune]node{},
									},
									'd': node{
										isWord: true,
										children: map[rune]node{
											's': node{
												isWord:   true,
												children: map[rune]node{},
											},
										},
									},
									't': node{
										isWord: false,
										children: map[rune]node{
											'h': node{
												isWord: false,
												children: map[rune]node{
													'l': node{
														isWord: false,
														children: map[rune]node{
															'e': node{
																isWord: false,
																children: map[rune]node{
																	's': node{
																		isWord: false,
																		children: map[rune]node{
																			's': node{
																				isWord:   true,
																				children: map[rune]node{},
																			},
																		},
																	},
																},
															},
														},
													},
													'f': node{
														isWord: false,
														children: map[rune]node{
															'u': node{
																isWord: false,
																children: map[rune]node{
																	'l': node{
																		isWord:   true,
																		children: map[rune]node{},
																	},
																},
															},
														},
													},
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}

	t.Run("Test adding one word", func(t *testing.T) {
		gd.addWord("word")
		gd.addWord("words")
		gd.addWord("work")
		gd.addWord("worthless")
		gd.addWord("worthful")

		if reflect.DeepEqual(gd.children['w'], wNode) != true {
			t.Errorf("Adding 5 words to graph has errors. Node: \n%v\n is different from example: \n%v\n", gd.children['w'], wNode)
			return
		}

	})
}
