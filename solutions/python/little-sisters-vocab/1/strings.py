"""Functions for creating, transforming, and adding prefixes to strings."""


def add_prefix_un(word):
    return f"un{word}"


def make_word_groups(vocab_words):
    return f' :: {vocab_words[0]}'.join(vocab_words)


def remove_suffix_ness(word):
    """Remove the suffix from the word while keeping spelling in mind.

    :param word: str - of word to remove suffix from.
    :return: str - of word with suffix removed & spelling adjusted.

    For example: "heaviness" becomes "heavy", but "sadness" becomes "sad".
    """

    word = word[:len(word)-4]
    if word[-1] == 'i':
        word = word[:-1] + 'y'
    return word


def adjective_to_verb(sentence, index):
    """Change the adjective within the sentence to a verb.

    :param sentence: str - that uses the word in sentence.
    :param index: int - index of the word to remove and transform.
    :return: str - word that changes the extracted adjective to a verb.

    For example, ("It got dark as the sun set.", 2) becomes "darken".
    """

    return sentence.split()[index].replace('.','')+'en'
