
---
layout: "../../layouts/LayoutSingle.astro"
title: cs121-quiz1
---
	
# Quiz 1

1. Information retrival
    - Field concerned with the structure analysis organization storage searching and retrieval of infromation
2. DB vs Search Engines
"find accounts whose balance is greater than 100"
- Well structured and will defined
- Not just text but numbers 
- Requires use of formal lang and logic 
3. Searching 
- Vague, informal/adhoc queries/ requires engine to infer meaning of words and sentences

# Retrieval Models
- Conceptual designs of waht to pay attention to when matchin a query and a document
- Good retrieval model: find docs relevant to the person who submitted the query
- basis for ranking algos

# Classic IR Assumptions
- Corpus: fixed doc collection
- Goal: retrieve info content relevant to the needed info

# IR Goal
- Relegance: query Q stored in document D -> exists score a R(Q,D)
- Maximize R(Q,D): context is ognored, user is ignored, corpus is static and centralized

# Web IR
- Web cant store any centralized memory
- Web needs to update constantl
- There is adversarial info to avoid
- One interface for hugely diverent needs 

# Web Search Engines 
- Must crawl tera to peta bytes of pages and provide subsecond repsonses o

# Text Transformation
- Parser = tokenzier + structure
- stopping: removing filler words
- stemming: grouping similar words
- link extraction and anaalysis: how popuilar a certain link is 
- info extraction: text structure
- classifier: topic, contentnon content, spam

# Index Creation
- Statistics
- Term weighting: how important a document is 
- Index inversion: doc -> term -> term -> doc
- Index distribution: essential for large indexes

# Ranking
- Relevance score: how well doc matches query
- Performance optimization: decrease response time
- Distribution: query broker, allocate queries to different processors

# Evaluation
- Loggin: improves search egine in the long run
- Ranking analysis: considering the user behavior, se analyzes if the ranking heuristics work well 
- Performance analysis: Huimans expect fast answers so the search engine is monitored to improve performance in the long run 


