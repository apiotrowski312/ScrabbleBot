package gaddag

import (
	"errors"
	"reflect"
	"testing"

	"github.com/bmizerany/assert"
)

func Test_GaddagAddWord_One(t *testing.T) {
	gd := Node{
		children: map[rune]Node{},
	}

	wNode := Node{
		isWord: false,
		children: map[rune]Node{
			'.': Node{
				isWord: false,
				children: map[rune]Node{
					'o': Node{
						isWord: false,
						children: map[rune]Node{
							'r': Node{
								isWord: false,
								children: map[rune]Node{
									'd': Node{
										isWord:   true,
										children: map[rune]Node{},
									},
								},
							},
						},
					},
				},
			},
		},
	}

	rNode := Node{
		isWord: false,
		children: map[rune]Node{
			'o': Node{
				isWord: false,
				children: map[rune]Node{
					'w': Node{
						isWord: false,
						children: map[rune]Node{
							'.': Node{
								isWord: false,
								children: map[rune]Node{
									'd': Node{
										isWord:   true,
										children: map[rune]Node{},
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

func Test_GaddagAddWords_Five(t *testing.T) {
	gd := Node{
		children: map[rune]Node{},
	}

	wNode := Node{
		isWord: false,
		children: map[rune]Node{
			'.': Node{
				isWord: false,
				children: map[rune]Node{
					'o': Node{
						isWord: false,
						children: map[rune]Node{
							'r': Node{
								isWord: false,
								children: map[rune]Node{
									'k': Node{
										isWord:   true,
										children: map[rune]Node{},
									},
									'd': Node{
										isWord: true,
										children: map[rune]Node{
											's': Node{
												isWord:   true,
												children: map[rune]Node{},
											},
										},
									},
									't': Node{
										isWord: false,
										children: map[rune]Node{
											'h': Node{
												isWord: false,
												children: map[rune]Node{
													'l': Node{
														isWord: false,
														children: map[rune]Node{
															'e': Node{
																isWord: false,
																children: map[rune]Node{
																	's': Node{
																		isWord: false,
																		children: map[rune]Node{
																			's': Node{
																				isWord:   true,
																				children: map[rune]Node{},
																			},
																		},
																	},
																},
															},
														},
													},
													'f': Node{
														isWord: false,
														children: map[rune]Node{
															'u': Node{
																isWord: false,
																children: map[rune]Node{
																	'l': Node{
																		isWord:   true,
																		children: map[rune]Node{},
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

	t.Run("Test adding five word", func(t *testing.T) {
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

func Test_CreateGraph(t *testing.T) {
	wNode := Node{
		isWord: false,
		children: map[rune]Node{
			'.': Node{
				isWord: false,
				children: map[rune]Node{
					'o': Node{
						isWord: false,
						children: map[rune]Node{
							'r': Node{
								isWord: false,
								children: map[rune]Node{
									'k': Node{
										isWord:   true,
										children: map[rune]Node{},
									},
									'd': Node{
										isWord: true,
										children: map[rune]Node{
											's': Node{
												isWord:   true,
												children: map[rune]Node{},
											},
										},
									},
									't': Node{
										isWord: false,
										children: map[rune]Node{
											'h': Node{
												isWord: false,
												children: map[rune]Node{
													'l': Node{
														isWord: false,
														children: map[rune]Node{
															'e': Node{
																isWord: false,
																children: map[rune]Node{
																	's': Node{
																		isWord: false,
																		children: map[rune]Node{
																			's': Node{
																				isWord:   true,
																				children: map[rune]Node{},
																			},
																		},
																	},
																},
															},
														},
													},
													'f': Node{
														isWord: false,
														children: map[rune]Node{
															'u': Node{
																isWord: false,
																children: map[rune]Node{
																	'l': Node{
																		isWord:   true,
																		children: map[rune]Node{},
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

	t.Run("Test creating graph", func(t *testing.T) {
		gaddagRoot, err := CreateGraph("../exampleData/tiny_english.txt")

		assert.Equal(t, err, nil, "There should be no error.")

		if reflect.DeepEqual(gaddagRoot.children['w'], wNode) != true {
			t.Errorf("Adding 5 words to graph has errors. Node: \n%v\n is different from example: \n%v\n", gaddagRoot.children['w'], wNode)
			return
		}

	})
}

func Test_IsWordValid(t *testing.T) {
	gaddagRoot, _ := CreateGraph("../exampleData/tiny_english.txt")

	isOk, err := gaddagRoot.IsWordValid("w.ord")
	assert.Equal(t, true, isOk)
	assert.Equal(t, nil, err)
	isOk, err = gaddagRoot.IsWordValid("w.ordlist")
	assert.Equal(t, false, isOk)
	assert.Equal(t, errors.New("Word w.ordlist is not in dictionary"), err)
	isOk, err = gaddagRoot.IsWordValid("w.orthless")
	assert.Equal(t, true, isOk)
	assert.Equal(t, nil, err)
	isOk, err = gaddagRoot.IsWordValid("w.orth")
	assert.Equal(t, false, isOk)
	assert.Equal(t, errors.New("Word w.orth is not in dictionary"), err)
	isOk, err = gaddagRoot.IsWordValid("ob.ss")
	assert.Equal(t, true, isOk)
	assert.Equal(t, nil, err)
}

func Benchmark_CreateGraph_5Words(b *testing.B) {
	for n := 0; n < b.N; n++ {
		CreateGraph("../exampleData/tiny_english.txt")
	}
}

func Benchmark_CreateGraph_2kWords(b *testing.B) {
	for n := 0; n < b.N; n++ {
		CreateGraph("../exampleData/2k_english.txt")
	}
}

func Benchmark_CreateGraph_20kWords(b *testing.B) {
	for n := 0; n < b.N; n++ {
		CreateGraph("../exampleData/20k_english.txt")
	}
}

func Benchmark_CreateGraph_280kWords(b *testing.B) {
	for n := 0; n < b.N; n++ {
		CreateGraph("../exampleData/collins_official_scrabble_2019.txt")
	}
}
