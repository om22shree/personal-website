# The 520-Problem DSA Curriculum

> **Purpose:** a deliberately sequenced curriculum for interview/OA preparation. This is not a LeetCode tag dump. Problems are ordered to *teach patterns*, then force transfer.

## How to use this file

Do **not** race through all 520 once. Your goal is recognition + reconstruction, not green-checkmark collection.

### The 3-pass rule

1. **Learn pass (same day)** — give a new problem ~15–25 focused minutes. If you have no viable direction, use a hint/editorial. Then close it and implement from a blank editor.
2. **Recall pass (+1 day)** — re-solve without notes. If you blank, write the pattern name and invariant first, then code.
3. **Retention pass (+7 days)** — timed re-solve. The problem is considered learned only when you can explain *why this pattern applies* before writing code.

### Module gate

Move to the next module when you can do both:

- identify the likely pattern on **at least 8/10 random problems** from the module, and
- solve **5 randomly selected problems** from the module without looking at prior code.

### When your brain goes blank

Ask in this order:

1. What is the input structure: array/string/list/tree/graph?
2. What special property exists: sorted, contiguous, unique, bounded, acyclic, weighted?
3. What is the output asking for: existence, count, optimum, construction, mutation?
4. What brute-force solution is obviously correct?
5. Which repeated work in that brute force can a data structure/pattern eliminate?

### Stage legend

- **F — Foundation:** learn the primitive.
- **B — Build:** same primitive with a meaningful twist.
- **T — Transfer:** combine ideas or handle a less-obvious formulation.

---

## Curriculum map

| Module | Topic | Problems |
|---:|---|---:|
| 01 | Arrays & Strings Fundamentals | 20 |
| 02 | Hash Maps, Sets & Frequency Tables | 20 |
| 03 | Prefix Sums, Difference Arrays & Counting | 20 |
| 04 | Two Pointers | 20 |
| 05 | Sliding Window | 20 |
| 06 | Sorting, Buckets & Ordering | 20 |
| 07 | Binary Search | 20 |
| 08 | Intervals & Sweep-Line Thinking | 20 |
| 09 | Stack Fundamentals & Parsing | 20 |
| 10 | Monotonic Stack | 20 |
| 11 | Queues, Deques & Streaming Windows | 20 |
| 12 | Linked Lists | 20 |
| 13 | Recursion & Backtracking | 20 |
| 14 | Binary Trees & BST Fundamentals | 20 |
| 15 | Advanced Trees, Paths & Reconstruction | 20 |
| 16 | Heaps & Priority Queues | 20 |
| 17 | Greedy | 20 |
| 18 | Graphs: BFS & DFS | 20 |
| 19 | Union-Find, DAGs & Topological Sort | 20 |
| 20 | Shortest Paths, Weighted Graphs & MST | 20 |
| 21 | Dynamic Programming I: 1D State | 20 |
| 22 | Dynamic Programming II: Grids & Knapsack | 20 |
| 23 | Dynamic Programming III: Strings, Sequences & Games | 20 |
| 24 | Tries & Bit Manipulation | 20 |
| 25 | Math, Geometry & Matrix Manipulation | 20 |
| 26 | Design Problems & OA Hardening | 20 |

**Total: 520 unique problems.**

---

## 01 — Arrays & Strings Fundamentals

**Goal:** Become comfortable reading constraints, traversing arrays/strings, mutating in place, and separating physical storage from logical length.

**Recognition cue:** *What is the simplest one-pass state I can maintain? Do I need mutation, a new output, or just an index?*

- [ ] **001. [Running Sum of 1d Array](https://leetcode.com/problems/running-sum-of-1d-array/)** — `F`
- [ ] **002. [Richest Customer Wealth](https://leetcode.com/problems/richest-customer-wealth/)** — `F`
- [ ] **003. [Build Array from Permutation](https://leetcode.com/problems/build-array-from-permutation/)** — `F`
- [ ] **004. [Concatenation of Array](https://leetcode.com/problems/concatenation-of-array/)** — `F`
- [ ] **005. [Shuffle the Array](https://leetcode.com/problems/shuffle-the-array/)** — `F`
- [ ] **006. [Kids With the Greatest Number of Candies](https://leetcode.com/problems/kids-with-the-greatest-number-of-candies/)** — `F`
- [ ] **007. [Final Value of Variable After Performing Operations](https://leetcode.com/problems/final-value-of-variable-after-performing-operations/)** — `B`
- [ ] **008. [Maximum Number of Words Found in Sentences](https://leetcode.com/problems/maximum-number-of-words-found-in-sentences/)** — `B`
- [ ] **009. [Move Zeroes](https://leetcode.com/problems/move-zeroes/)** — `B`
- [ ] **010. [Remove Element](https://leetcode.com/problems/remove-element/)** — `B`
- [ ] **011. [Remove Duplicates from Sorted Array](https://leetcode.com/problems/remove-duplicates-from-sorted-array/)** — `B`
- [ ] **012. [Plus One](https://leetcode.com/problems/plus-one/)** — `B`
- [ ] **013. [Merge Sorted Array](https://leetcode.com/problems/merge-sorted-array/)** — `B`
- [ ] **014. [Best Time to Buy and Sell Stock](https://leetcode.com/problems/best-time-to-buy-and-sell-stock/)** — `B`
- [ ] **015. [Maximum Subarray](https://leetcode.com/problems/maximum-subarray/)** — `T`
- [ ] **016. [Majority Element](https://leetcode.com/problems/majority-element/)** — `T`
- [ ] **017. [Rotate Array](https://leetcode.com/problems/rotate-array/)** — `T`
- [ ] **018. [Product of Array Except Self](https://leetcode.com/problems/product-of-array-except-self/)** — `T`
- [ ] **019. [Spiral Matrix](https://leetcode.com/problems/spiral-matrix/)** — `T`
- [ ] **020. [Set Matrix Zeroes](https://leetcode.com/problems/set-matrix-zeroes/)** — `T`

**Exit check:** pick 3 problems from Foundation, 2 from Build, and 1 from Transfer at random. Before coding each, state the invariant/pattern in one sentence.

---

## 02 — Hash Maps, Sets & Frequency Tables

**Goal:** Make frequency tables, membership tests, and key→state mappings automatic.

**Recognition cue:** *Am I repeatedly asking 'have I seen X?', 'how many X?', or 'what state belongs to X?'*

- [ ] **021. [Two Sum](https://leetcode.com/problems/two-sum/)** — `F`
- [ ] **022. [Contains Duplicate](https://leetcode.com/problems/contains-duplicate/)** — `F`
- [ ] **023. [Valid Anagram](https://leetcode.com/problems/valid-anagram/)** — `F`
- [ ] **024. [Ransom Note](https://leetcode.com/problems/ransom-note/)** — `F`
- [ ] **025. [First Unique Character in a String](https://leetcode.com/problems/first-unique-character-in-a-string/)** — `F`
- [ ] **026. [Isomorphic Strings](https://leetcode.com/problems/isomorphic-strings/)** — `F`
- [ ] **027. [Word Pattern](https://leetcode.com/problems/word-pattern/)** — `B`
- [ ] **028. [Intersection of Two Arrays](https://leetcode.com/problems/intersection-of-two-arrays/)** — `B`
- [ ] **029. [Happy Number](https://leetcode.com/problems/happy-number/)** — `B`
- [ ] **030. [Group Anagrams](https://leetcode.com/problems/group-anagrams/)** — `B`
- [ ] **031. [Top K Frequent Elements](https://leetcode.com/problems/top-k-frequent-elements/)** — `B`
- [ ] **032. [Longest Consecutive Sequence](https://leetcode.com/problems/longest-consecutive-sequence/)** — `B`
- [ ] **033. [4Sum II](https://leetcode.com/problems/4sum-ii/)** — `B`
- [ ] **034. [Sort Characters By Frequency](https://leetcode.com/problems/sort-characters-by-frequency/)** — `B`
- [ ] **035. [Determine if Two Strings Are Close](https://leetcode.com/problems/determine-if-two-strings-are-close/)** — `T`
- [ ] **036. [Unique Number of Occurrences](https://leetcode.com/problems/unique-number-of-occurrences/)** — `T`
- [ ] **037. [Find Common Characters](https://leetcode.com/problems/find-common-characters/)** — `T`
- [ ] **038. [Jewels and Stones](https://leetcode.com/problems/jewels-and-stones/)** — `T`
- [ ] **039. [Find the Difference](https://leetcode.com/problems/find-the-difference/)** — `T`
- [ ] **040. [Longest Palindrome](https://leetcode.com/problems/longest-palindrome/)** — `T`

**Exit check:** pick 3 problems from Foundation, 2 from Build, and 1 from Transfer at random. Before coding each, state the invariant/pattern in one sentence.

---

## 03 — Prefix Sums, Difference Arrays & Counting

**Goal:** Learn to turn repeated range work into cumulative state and to recognize difference-array updates.

**Recognition cue:** *Can I precompute cumulative information so a range/query becomes O(1)?*

- [ ] **041. [Find Pivot Index](https://leetcode.com/problems/find-pivot-index/)** — `F`
- [ ] **042. [Range Sum Query - Immutable](https://leetcode.com/problems/range-sum-query-immutable/)** — `F`
- [ ] **043. [Left and Right Sum Differences](https://leetcode.com/problems/left-and-right-sum-differences/)** — `F`
- [ ] **044. [Find the Highest Altitude](https://leetcode.com/problems/find-the-highest-altitude/)** — `F`
- [ ] **045. [Minimum Value to Get Positive Step by Step Sum](https://leetcode.com/problems/minimum-value-to-get-positive-step-by-step-sum/)** — `F`
- [ ] **046. [Continuous Subarray Sum](https://leetcode.com/problems/continuous-subarray-sum/)** — `F`
- [ ] **047. [Contiguous Array](https://leetcode.com/problems/contiguous-array/)** — `B`
- [ ] **048. [Binary Subarrays With Sum](https://leetcode.com/problems/binary-subarrays-with-sum/)** — `B`
- [ ] **049. [Corporate Flight Bookings](https://leetcode.com/problems/corporate-flight-bookings/)** — `B`
- [ ] **050. [Range Addition](https://leetcode.com/problems/range-addition/)** — `B`
- [ ] **051. [Matrix Block Sum](https://leetcode.com/problems/matrix-block-sum/)** — `B`
- [ ] **052. [Range Sum Query 2D - Immutable](https://leetcode.com/problems/range-sum-query-2d-immutable/)** — `B`
- [ ] **053. [Number of Ways to Split Array](https://leetcode.com/problems/number-of-ways-to-split-array/)** — `B`
- [ ] **054. [Sum of Absolute Differences in a Sorted Array](https://leetcode.com/problems/sum-of-absolute-differences-in-a-sorted-array/)** — `B`
- [ ] **055. [Maximum Population Year](https://leetcode.com/problems/maximum-population-year/)** — `T`
- [ ] **056. [Plates Between Candles](https://leetcode.com/problems/plates-between-candles/)** — `T`
- [ ] **057. [Ways to Make a Fair Array](https://leetcode.com/problems/ways-to-make-a-fair-array/)** — `T`
- [ ] **058. [K Radius Subarray Averages](https://leetcode.com/problems/k-radius-subarray-averages/)** — `T`
- [ ] **059. [XOR Queries of a Subarray](https://leetcode.com/problems/xor-queries-of-a-subarray/)** — `T`
- [ ] **060. [Shifting Letters II](https://leetcode.com/problems/shifting-letters-ii/)** — `T`

**Exit check:** pick 3 problems from Foundation, 2 from Build, and 1 from Transfer at random. Before coding each, state the invariant/pattern in one sentence.

---

## 04 — Two Pointers

**Goal:** Internalize inward, same-direction, and partition-style pointer movement.

**Recognition cue:** *What invariant lets me discard one side or move one pointer without losing an answer?*

- [ ] **061. [Valid Palindrome](https://leetcode.com/problems/valid-palindrome/)** — `F`
- [ ] **062. [Reverse String](https://leetcode.com/problems/reverse-string/)** — `F`
- [ ] **063. [Squares of a Sorted Array](https://leetcode.com/problems/squares-of-a-sorted-array/)** — `F`
- [ ] **064. [Two Sum II - Input Array Is Sorted](https://leetcode.com/problems/two-sum-ii-input-array-is-sorted/)** — `F`
- [ ] **065. [Remove Duplicates from Sorted Array II](https://leetcode.com/problems/remove-duplicates-from-sorted-array-ii/)** — `F`
- [ ] **066. [Container With Most Water](https://leetcode.com/problems/container-with-most-water/)** — `F`
- [ ] **067. [3Sum](https://leetcode.com/problems/3sum/)** — `B`
- [ ] **068. [3Sum Closest](https://leetcode.com/problems/3sum-closest/)** — `B`
- [ ] **069. [4Sum](https://leetcode.com/problems/4sum/)** — `B`
- [ ] **070. [Sort Colors](https://leetcode.com/problems/sort-colors/)** — `B`
- [ ] **071. [Reverse Vowels of a String](https://leetcode.com/problems/reverse-vowels-of-a-string/)** — `B`
- [ ] **072. [Valid Palindrome II](https://leetcode.com/problems/valid-palindrome-ii/)** — `B`
- [ ] **073. [String Compression](https://leetcode.com/problems/string-compression/)** — `B`
- [ ] **074. [Merge Strings Alternately](https://leetcode.com/problems/merge-strings-alternately/)** — `B`
- [ ] **075. [Append Characters to String to Make Subsequence](https://leetcode.com/problems/append-characters-to-string-to-make-subsequence/)** — `T`
- [ ] **076. [Minimum Length of String After Deleting Similar Ends](https://leetcode.com/problems/minimum-length-of-string-after-deleting-similar-ends/)** — `T`
- [ ] **077. [Boats to Save People](https://leetcode.com/problems/boats-to-save-people/)** — `T`
- [ ] **078. [Trapping Rain Water](https://leetcode.com/problems/trapping-rain-water/)** — `T`
- [ ] **079. [Minimum Number of Moves to Make Palindrome](https://leetcode.com/problems/minimum-number-of-moves-to-make-palindrome/)** — `T`
- [ ] **080. [Next Permutation](https://leetcode.com/problems/next-permutation/)** — `T`

**Exit check:** pick 3 problems from Foundation, 2 from Build, and 1 from Transfer at random. Before coding each, state the invariant/pattern in one sentence.

---

## 05 — Sliding Window

**Goal:** Recognize contiguous-range problems and maintain a valid window incrementally.

**Recognition cue:** *Can I expand right, then move left only when an invariant breaks?*

- [ ] **081. [Maximum Average Subarray I](https://leetcode.com/problems/maximum-average-subarray-i/)** — `F`
- [ ] **082. [Contains Duplicate II](https://leetcode.com/problems/contains-duplicate-ii/)** — `F`
- [ ] **083. [Longest Substring Without Repeating Characters](https://leetcode.com/problems/longest-substring-without-repeating-characters/)** — `F`
- [ ] **084. [Minimum Size Subarray Sum](https://leetcode.com/problems/minimum-size-subarray-sum/)** — `F`
- [ ] **085. [Longest Repeating Character Replacement](https://leetcode.com/problems/longest-repeating-character-replacement/)** — `F`
- [ ] **086. [Permutation in String](https://leetcode.com/problems/permutation-in-string/)** — `F`
- [ ] **087. [Max Consecutive Ones III](https://leetcode.com/problems/max-consecutive-ones-iii/)** — `B`
- [ ] **088. [Fruit Into Baskets](https://leetcode.com/problems/fruit-into-baskets/)** — `B`
- [ ] **089. [Subarray Product Less Than K](https://leetcode.com/problems/subarray-product-less-than-k/)** — `B`
- [ ] **090. [Grumpy Bookstore Owner](https://leetcode.com/problems/grumpy-bookstore-owner/)** — `B`
- [ ] **091. [Frequency of the Most Frequent Element](https://leetcode.com/problems/frequency-of-the-most-frequent-element/)** — `B`
- [ ] **092. [Minimum Window Substring](https://leetcode.com/problems/minimum-window-substring/)** — `B`
- [ ] **093. [Sliding Window Maximum](https://leetcode.com/problems/sliding-window-maximum/)** — `B`
- [ ] **094. [Longest Subarray of 1's After Deleting One Element](https://leetcode.com/problems/longest-subarray-of-1s-after-deleting-one-element/)** — `B`
- [ ] **095. [Get Equal Substrings Within Budget](https://leetcode.com/problems/get-equal-substrings-within-budget/)** — `T`
- [ ] **096. [Number of Substrings Containing All Three Characters](https://leetcode.com/problems/number-of-substrings-containing-all-three-characters/)** — `T`
- [ ] **097. [Count Number of Nice Subarrays](https://leetcode.com/problems/count-number-of-nice-subarrays/)** — `T`
- [ ] **098. [Replace the Substring for Balanced String](https://leetcode.com/problems/replace-the-substring-for-balanced-string/)** — `T`
- [ ] **099. [Maximum Points You Can Obtain from Cards](https://leetcode.com/problems/maximum-points-you-can-obtain-from-cards/)** — `T`
- [ ] **100. [Take K of Each Character From Left and Right](https://leetcode.com/problems/take-k-of-each-character-from-left-and-right/)** — `T`

**Exit check:** pick 3 problems from Foundation, 2 from Build, and 1 from Transfer at random. Before coding each, state the invariant/pattern in one sentence.

---

## 06 — Sorting, Buckets & Ordering

**Goal:** Use ordering as a tool: sorting, custom keys, buckets, ranks, and order statistics.

**Recognition cue:** *Would sorting make the structure obvious enough that the remaining scan is trivial?*

- [ ] **101. [Sort Array By Parity](https://leetcode.com/problems/sort-array-by-parity/)** — `F`
- [ ] **102. [Sort Array by Increasing Frequency](https://leetcode.com/problems/sort-array-by-increasing-frequency/)** — `F`
- [ ] **103. [Relative Sort Array](https://leetcode.com/problems/relative-sort-array/)** — `F`
- [ ] **104. [Height Checker](https://leetcode.com/problems/height-checker/)** — `F`
- [ ] **105. [Array Partition](https://leetcode.com/problems/array-partition/)** — `F`
- [ ] **106. [Minimum Number of Moves to Seat Everyone](https://leetcode.com/problems/minimum-number-of-moves-to-seat-everyone/)** — `F`
- [ ] **107. [Minimum Difference Between Highest and Lowest of K Scores](https://leetcode.com/problems/minimum-difference-between-highest-and-lowest-of-k-scores/)** — `B`
- [ ] **108. [Maximum Product Difference Between Two Pairs](https://leetcode.com/problems/maximum-product-difference-between-two-pairs/)** — `B`
- [ ] **109. [Relative Ranks](https://leetcode.com/problems/relative-ranks/)** — `B`
- [ ] **110. [Third Maximum Number](https://leetcode.com/problems/third-maximum-number/)** — `B`
- [ ] **111. [Custom Sort String](https://leetcode.com/problems/custom-sort-string/)** — `B`
- [ ] **112. [Sort Integers by The Number of 1 Bits](https://leetcode.com/problems/sort-integers-by-the-number-of-1-bits/)** — `B`
- [ ] **113. [Rank Transform of an Array](https://leetcode.com/problems/rank-transform-of-an-array/)** — `B`
- [ ] **114. [Sort the People](https://leetcode.com/problems/sort-the-people/)** — `B`
- [ ] **115. [H-Index](https://leetcode.com/problems/h-index/)** — `T`
- [ ] **116. [Largest Number](https://leetcode.com/problems/largest-number/)** — `T`
- [ ] **117. [Sort an Array](https://leetcode.com/problems/sort-an-array/)** — `T`
- [ ] **118. [Wiggle Sort II](https://leetcode.com/problems/wiggle-sort-ii/)** — `T`
- [ ] **119. [Pancake Sorting](https://leetcode.com/problems/pancake-sorting/)** — `T`
- [ ] **120. [Maximum Gap](https://leetcode.com/problems/maximum-gap/)** — `T`

**Exit check:** pick 3 problems from Foundation, 2 from Build, and 1 from Transfer at random. Before coding each, state the invariant/pattern in one sentence.

---

## 07 — Binary Search

**Goal:** Learn both classic binary search and binary search on the answer.

**Recognition cue:** *Is there a monotonic yes/no predicate over an ordered search space?*

- [ ] **121. [Binary Search](https://leetcode.com/problems/binary-search/)** — `F`
- [ ] **122. [Search Insert Position](https://leetcode.com/problems/search-insert-position/)** — `F`
- [ ] **123. [Guess Number Higher or Lower](https://leetcode.com/problems/guess-number-higher-or-lower/)** — `F`
- [ ] **124. [First Bad Version](https://leetcode.com/problems/first-bad-version/)** — `F`
- [ ] **125. [Sqrt(x)](https://leetcode.com/problems/sqrtx/)** — `F`
- [ ] **126. [Valid Perfect Square](https://leetcode.com/problems/valid-perfect-square/)** — `F`
- [ ] **127. [Find First and Last Position of Element in Sorted Array](https://leetcode.com/problems/find-first-and-last-position-of-element-in-sorted-array/)** — `B`
- [ ] **128. [Search in Rotated Sorted Array](https://leetcode.com/problems/search-in-rotated-sorted-array/)** — `B`
- [ ] **129. [Search in Rotated Sorted Array II](https://leetcode.com/problems/search-in-rotated-sorted-array-ii/)** — `B`
- [ ] **130. [Find Minimum in Rotated Sorted Array](https://leetcode.com/problems/find-minimum-in-rotated-sorted-array/)** — `B`
- [ ] **131. [Find Peak Element](https://leetcode.com/problems/find-peak-element/)** — `B`
- [ ] **132. [Peak Index in a Mountain Array](https://leetcode.com/problems/peak-index-in-a-mountain-array/)** — `B`
- [ ] **133. [Find K Closest Elements](https://leetcode.com/problems/find-k-closest-elements/)** — `B`
- [ ] **134. [Search a 2D Matrix](https://leetcode.com/problems/search-a-2d-matrix/)** — `B`
- [ ] **135. [Koko Eating Bananas](https://leetcode.com/problems/koko-eating-bananas/)** — `T`
- [ ] **136. [Capacity To Ship Packages Within D Days](https://leetcode.com/problems/capacity-to-ship-packages-within-d-days/)** — `T`
- [ ] **137. [Minimum Speed to Arrive on Time](https://leetcode.com/problems/minimum-speed-to-arrive-on-time/)** — `T`
- [ ] **138. [Split Array Largest Sum](https://leetcode.com/problems/split-array-largest-sum/)** — `T`
- [ ] **139. [Median of Two Sorted Arrays](https://leetcode.com/problems/median-of-two-sorted-arrays/)** — `T`
- [ ] **140. [Time Based Key-Value Store](https://leetcode.com/problems/time-based-key-value-store/)** — `T`

**Exit check:** pick 3 problems from Foundation, 2 from Build, and 1 from Transfer at random. Before coding each, state the invariant/pattern in one sentence.

---

## 08 — Intervals & Sweep-Line Thinking

**Goal:** Reason about starts, ends, overlap, event ordering, and active intervals.

**Recognition cue:** *After sorting by start/end, what information about the previous/active intervals is sufficient?*

- [ ] **141. [Summary Ranges](https://leetcode.com/problems/summary-ranges/)** — `F`
- [ ] **142. [Merge Intervals](https://leetcode.com/problems/merge-intervals/)** — `F`
- [ ] **143. [Insert Interval](https://leetcode.com/problems/insert-interval/)** — `F`
- [ ] **144. [Meeting Rooms](https://leetcode.com/problems/meeting-rooms/)** — `F`
- [ ] **145. [Meeting Rooms II](https://leetcode.com/problems/meeting-rooms-ii/)** — `F`
- [ ] **146. [Non-overlapping Intervals](https://leetcode.com/problems/non-overlapping-intervals/)** — `F`
- [ ] **147. [Minimum Number of Arrows to Burst Balloons](https://leetcode.com/problems/minimum-number-of-arrows-to-burst-balloons/)** — `B`
- [ ] **148. [Interval List Intersections](https://leetcode.com/problems/interval-list-intersections/)** — `B`
- [ ] **149. [Remove Covered Intervals](https://leetcode.com/problems/remove-covered-intervals/)** — `B`
- [ ] **150. [Meeting Scheduler](https://leetcode.com/problems/meeting-scheduler/)** — `B`
- [ ] **151. [My Calendar I](https://leetcode.com/problems/my-calendar-i/)** — `B`
- [ ] **152. [My Calendar II](https://leetcode.com/problems/my-calendar-ii/)** — `B`
- [ ] **153. [Employee Free Time](https://leetcode.com/problems/employee-free-time/)** — `B`
- [ ] **154. [Video Stitching](https://leetcode.com/problems/video-stitching/)** — `B`
- [ ] **155. [Minimum Interval to Include Each Query](https://leetcode.com/problems/minimum-interval-to-include-each-query/)** — `T`
- [ ] **156. [Data Stream as Disjoint Intervals](https://leetcode.com/problems/data-stream-as-disjoint-intervals/)** — `T`
- [ ] **157. [Amount of New Area Painted Each Day](https://leetcode.com/problems/amount-of-new-area-painted-each-day/)** — `T`
- [ ] **158. [Divide Intervals Into Minimum Number of Groups](https://leetcode.com/problems/divide-intervals-into-minimum-number-of-groups/)** — `T`
- [ ] **159. [Points That Intersect With Cars](https://leetcode.com/problems/points-that-intersect-with-cars/)** — `T`
- [ ] **160. [Maximum Number of Events That Can Be Attended](https://leetcode.com/problems/maximum-number-of-events-that-can-be-attended/)** — `T`

**Exit check:** pick 3 problems from Foundation, 2 from Build, and 1 from Transfer at random. Before coding each, state the invariant/pattern in one sentence.

---

## 09 — Stack Fundamentals & Parsing

**Goal:** Use stacks for nested structure, undo-like behavior, parsing, and deferred work.

**Recognition cue:** *Do I need the most recent unmatched/unresolved item first?*

- [ ] **161. [Valid Parentheses](https://leetcode.com/problems/valid-parentheses/)** — `F`
- [ ] **162. [Baseball Game](https://leetcode.com/problems/baseball-game/)** — `F`
- [ ] **163. [Remove All Adjacent Duplicates In String](https://leetcode.com/problems/remove-all-adjacent-duplicates-in-string/)** — `F`
- [ ] **164. [Make The String Great](https://leetcode.com/problems/make-the-string-great/)** — `F`
- [ ] **165. [Build an Array With Stack Operations](https://leetcode.com/problems/build-an-array-with-stack-operations/)** — `F`
- [ ] **166. [Min Stack](https://leetcode.com/problems/min-stack/)** — `F`
- [ ] **167. [Evaluate Reverse Polish Notation](https://leetcode.com/problems/evaluate-reverse-polish-notation/)** — `B`
- [ ] **168. [Simplify Path](https://leetcode.com/problems/simplify-path/)** — `B`
- [ ] **169. [Score of Parentheses](https://leetcode.com/problems/score-of-parentheses/)** — `B`
- [ ] **170. [Decode String](https://leetcode.com/problems/decode-string/)** — `B`
- [ ] **171. [Asteroid Collision](https://leetcode.com/problems/asteroid-collision/)** — `B`
- [ ] **172. [Validate Stack Sequences](https://leetcode.com/problems/validate-stack-sequences/)** — `B`
- [ ] **173. [Removing Stars From a String](https://leetcode.com/problems/removing-stars-from-a-string/)** — `B`
- [ ] **174. [Minimum Remove to Make Valid Parentheses](https://leetcode.com/problems/minimum-remove-to-make-valid-parentheses/)** — `B`
- [ ] **175. [Basic Calculator II](https://leetcode.com/problems/basic-calculator-ii/)** — `T`
- [ ] **176. [Basic Calculator](https://leetcode.com/problems/basic-calculator/)** — `T`
- [ ] **177. [Exclusive Time of Functions](https://leetcode.com/problems/exclusive-time-of-functions/)** — `T`
- [ ] **178. [Longest Valid Parentheses](https://leetcode.com/problems/longest-valid-parentheses/)** — `T`
- [ ] **179. [Check If Word Is Valid After Substitutions](https://leetcode.com/problems/check-if-word-is-valid-after-substitutions/)** — `T`
- [ ] **180. [Parse Lisp Expression](https://leetcode.com/problems/parse-lisp-expression/)** — `T`

**Exit check:** pick 3 problems from Foundation, 2 from Build, and 1 from Transfer at random. Before coding each, state the invariant/pattern in one sentence.

---

## 10 — Monotonic Stack

**Goal:** Recognize nearest-greater/smaller patterns and maintain a monotonic frontier.

**Recognition cue:** *Can an element be permanently discarded once a stronger candidate arrives?*

- [ ] **181. [Next Greater Element I](https://leetcode.com/problems/next-greater-element-i/)** — `F`
- [ ] **182. [Next Greater Element II](https://leetcode.com/problems/next-greater-element-ii/)** — `F`
- [ ] **183. [Daily Temperatures](https://leetcode.com/problems/daily-temperatures/)** — `F`
- [ ] **184. [Final Prices With a Special Discount in a Shop](https://leetcode.com/problems/final-prices-with-a-special-discount-in-a-shop/)** — `F`
- [ ] **185. [Online Stock Span](https://leetcode.com/problems/online-stock-span/)** — `F`
- [ ] **186. [Next Greater Node In Linked List](https://leetcode.com/problems/next-greater-node-in-linked-list/)** — `F`
- [ ] **187. [132 Pattern](https://leetcode.com/problems/132-pattern/)** — `B`
- [ ] **188. [Remove K Digits](https://leetcode.com/problems/remove-k-digits/)** — `B`
- [ ] **189. [Car Fleet](https://leetcode.com/problems/car-fleet/)** — `B`
- [ ] **190. [Shortest Unsorted Continuous Subarray](https://leetcode.com/problems/shortest-unsorted-continuous-subarray/)** — `B`
- [ ] **191. [Maximum Width Ramp](https://leetcode.com/problems/maximum-width-ramp/)** — `B`
- [ ] **192. [Largest Rectangle in Histogram](https://leetcode.com/problems/largest-rectangle-in-histogram/)** — `B`
- [ ] **193. [Maximal Rectangle](https://leetcode.com/problems/maximal-rectangle/)** — `B`
- [ ] **194. [Sum of Subarray Minimums](https://leetcode.com/problems/sum-of-subarray-minimums/)** — `B`
- [ ] **195. [Sum of Subarray Ranges](https://leetcode.com/problems/sum-of-subarray-ranges/)** — `T`
- [ ] **196. [Number of Visible People in a Queue](https://leetcode.com/problems/number-of-visible-people-in-a-queue/)** — `T`
- [ ] **197. [Steps to Make Array Non-decreasing](https://leetcode.com/problems/steps-to-make-array-non-decreasing/)** — `T`
- [ ] **198. [Minimum Cost Tree From Leaf Values](https://leetcode.com/problems/minimum-cost-tree-from-leaf-values/)** — `T`
- [ ] **199. [Beautiful Towers II](https://leetcode.com/problems/beautiful-towers-ii/)** — `T`
- [ ] **200. [Total Strength of Wizards](https://leetcode.com/problems/total-strength-of-wizards/)** — `T`

**Exit check:** pick 3 problems from Foundation, 2 from Build, and 1 from Transfer at random. Before coding each, state the invariant/pattern in one sentence.

---

## 11 — Queues, Deques & Streaming Windows

**Goal:** Build intuition for FIFO processing, deque endpoints, and streaming/window extrema.

**Recognition cue:** *Do I need oldest-first processing, or candidates that expire from the front?*

- [ ] **201. [Implement Queue using Stacks](https://leetcode.com/problems/implement-queue-using-stacks/)** — `F`
- [ ] **202. [Implement Stack using Queues](https://leetcode.com/problems/implement-stack-using-queues/)** — `F`
- [ ] **203. [Number of Recent Calls](https://leetcode.com/problems/number-of-recent-calls/)** — `F`
- [ ] **204. [Design Circular Queue](https://leetcode.com/problems/design-circular-queue/)** — `F`
- [ ] **205. [Moving Average from Data Stream](https://leetcode.com/problems/moving-average-from-data-stream/)** — `F`
- [ ] **206. [Number of Students Unable to Eat Lunch](https://leetcode.com/problems/number-of-students-unable-to-eat-lunch/)** — `F`
- [ ] **207. [Time Needed to Buy Tickets](https://leetcode.com/problems/time-needed-to-buy-tickets/)** — `B`
- [ ] **208. [Find the Winner of the Circular Game](https://leetcode.com/problems/find-the-winner-of-the-circular-game/)** — `B`
- [ ] **209. [Dota2 Senate](https://leetcode.com/problems/dota2-senate/)** — `B`
- [ ] **210. [Reveal Cards In Increasing Order](https://leetcode.com/problems/reveal-cards-in-increasing-order/)** — `B`
- [ ] **211. [Design Front Middle Back Queue](https://leetcode.com/problems/design-front-middle-back-queue/)** — `B`
- [ ] **212. [First Unique Number](https://leetcode.com/problems/first-unique-number/)** — `B`
- [ ] **213. [Design Hit Counter](https://leetcode.com/problems/design-hit-counter/)** — `B`
- [ ] **214. [Longest Continuous Subarray With Absolute Diff Less Than or Equal to Limit](https://leetcode.com/problems/longest-continuous-subarray-with-absolute-diff-less-than-or-equal-to-limit/)** — `B`
- [ ] **215. [Continuous Subarrays](https://leetcode.com/problems/continuous-subarrays/)** — `T`
- [ ] **216. [Maximum Number of Robots Within Budget](https://leetcode.com/problems/maximum-number-of-robots-within-budget/)** — `T`
- [ ] **217. [Shortest Subarray with Sum at Least K](https://leetcode.com/problems/shortest-subarray-with-sum-at-least-k/)** — `T`
- [ ] **218. [Constrained Subsequence Sum](https://leetcode.com/problems/constrained-subsequence-sum/)** — `T`
- [ ] **219. [Jump Game VI](https://leetcode.com/problems/jump-game-vi/)** — `T`
- [ ] **220. [Max Value of Equation](https://leetcode.com/problems/max-value-of-equation/)** — `T`

**Exit check:** pick 3 problems from Foundation, 2 from Build, and 1 from Transfer at random. Before coding each, state the invariant/pattern in one sentence.

---

## 12 — Linked Lists

**Goal:** Make node identity, pointer rewiring, dummy heads, fast/slow pointers, and recursion on lists familiar.

**Recognition cue:** *Which exact node reference must survive this operation, and what pointer can I safely change?*

- [ ] **221. [Middle of the Linked List](https://leetcode.com/problems/middle-of-the-linked-list/)** — `F`
- [ ] **222. [Linked List Cycle](https://leetcode.com/problems/linked-list-cycle/)** — `F`
- [ ] **223. [Reverse Linked List](https://leetcode.com/problems/reverse-linked-list/)** — `F`
- [ ] **224. [Merge Two Sorted Lists](https://leetcode.com/problems/merge-two-sorted-lists/)** — `F`
- [ ] **225. [Remove Duplicates from Sorted List](https://leetcode.com/problems/remove-duplicates-from-sorted-list/)** — `F`
- [ ] **226. [Remove Linked List Elements](https://leetcode.com/problems/remove-linked-list-elements/)** — `F`
- [ ] **227. [Intersection of Two Linked Lists](https://leetcode.com/problems/intersection-of-two-linked-lists/)** — `B`
- [ ] **228. [Delete Node in a Linked List](https://leetcode.com/problems/delete-node-in-a-linked-list/)** — `B`
- [ ] **229. [Palindrome Linked List](https://leetcode.com/problems/palindrome-linked-list/)** — `B`
- [ ] **230. [Linked List Cycle II](https://leetcode.com/problems/linked-list-cycle-ii/)** — `B`
- [ ] **231. [Remove Nth Node From End of List](https://leetcode.com/problems/remove-nth-node-from-end-of-list/)** — `B`
- [ ] **232. [Swap Nodes in Pairs](https://leetcode.com/problems/swap-nodes-in-pairs/)** — `B`
- [ ] **233. [Add Two Numbers](https://leetcode.com/problems/add-two-numbers/)** — `B`
- [ ] **234. [Odd Even Linked List](https://leetcode.com/problems/odd-even-linked-list/)** — `B`
- [ ] **235. [Partition List](https://leetcode.com/problems/partition-list/)** — `T`
- [ ] **236. [Reorder List](https://leetcode.com/problems/reorder-list/)** — `T`
- [ ] **237. [Sort List](https://leetcode.com/problems/sort-list/)** — `T`
- [ ] **238. [Copy List with Random Pointer](https://leetcode.com/problems/copy-list-with-random-pointer/)** — `T`
- [ ] **239. [Flatten a Multilevel Doubly Linked List](https://leetcode.com/problems/flatten-a-multilevel-doubly-linked-list/)** — `T`
- [ ] **240. [Reverse Nodes in k-Group](https://leetcode.com/problems/reverse-nodes-in-k-group/)** — `T`

**Exit check:** pick 3 problems from Foundation, 2 from Build, and 1 from Transfer at random. Before coding each, state the invariant/pattern in one sentence.

---

## 13 — Recursion & Backtracking

**Goal:** Learn recursive state trees: choose → recurse → undo, with clear base cases.

**Recognition cue:** *What is my state, what choices exist now, and what makes a branch complete or invalid?*

- [ ] **241. [Fibonacci Number](https://leetcode.com/problems/fibonacci-number/)** — `F`
- [ ] **242. [Pow(x, n)](https://leetcode.com/problems/powx-n/)** — `F`
- [ ] **243. [Letter Combinations of a Phone Number](https://leetcode.com/problems/letter-combinations-of-a-phone-number/)** — `F`
- [ ] **244. [Generate Parentheses](https://leetcode.com/problems/generate-parentheses/)** — `F`
- [ ] **245. [Subsets](https://leetcode.com/problems/subsets/)** — `F`
- [ ] **246. [Subsets II](https://leetcode.com/problems/subsets-ii/)** — `F`
- [ ] **247. [Permutations](https://leetcode.com/problems/permutations/)** — `B`
- [ ] **248. [Permutations II](https://leetcode.com/problems/permutations-ii/)** — `B`
- [ ] **249. [Combinations](https://leetcode.com/problems/combinations/)** — `B`
- [ ] **250. [Combination Sum](https://leetcode.com/problems/combination-sum/)** — `B`
- [ ] **251. [Combination Sum II](https://leetcode.com/problems/combination-sum-ii/)** — `B`
- [ ] **252. [Combination Sum III](https://leetcode.com/problems/combination-sum-iii/)** — `B`
- [ ] **253. [Palindrome Partitioning](https://leetcode.com/problems/palindrome-partitioning/)** — `B`
- [ ] **254. [Restore IP Addresses](https://leetcode.com/problems/restore-ip-addresses/)** — `B`
- [ ] **255. [Word Search](https://leetcode.com/problems/word-search/)** — `T`
- [ ] **256. [Beautiful Arrangement](https://leetcode.com/problems/beautiful-arrangement/)** — `T`
- [ ] **257. [Matchsticks to Square](https://leetcode.com/problems/matchsticks-to-square/)** — `T`
- [ ] **258. [N-Queens](https://leetcode.com/problems/n-queens/)** — `T`
- [ ] **259. [N-Queens II](https://leetcode.com/problems/n-queens-ii/)** — `T`
- [ ] **260. [Sudoku Solver](https://leetcode.com/problems/sudoku-solver/)** — `T`

**Exit check:** pick 3 problems from Foundation, 2 from Build, and 1 from Transfer at random. Before coding each, state the invariant/pattern in one sentence.

---

## 14 — Binary Trees & BST Fundamentals

**Goal:** Make DFS/BFS tree traversal and basic BST invariants automatic.

**Recognition cue:** *What does the recursive call promise me about the subtree it returns/processes?*

- [ ] **261. [Maximum Depth of Binary Tree](https://leetcode.com/problems/maximum-depth-of-binary-tree/)** — `F`
- [ ] **262. [Same Tree](https://leetcode.com/problems/same-tree/)** — `F`
- [ ] **263. [Invert Binary Tree](https://leetcode.com/problems/invert-binary-tree/)** — `F`
- [ ] **264. [Symmetric Tree](https://leetcode.com/problems/symmetric-tree/)** — `F`
- [ ] **265. [Minimum Depth of Binary Tree](https://leetcode.com/problems/minimum-depth-of-binary-tree/)** — `F`
- [ ] **266. [Balanced Binary Tree](https://leetcode.com/problems/balanced-binary-tree/)** — `F`
- [ ] **267. [Diameter of Binary Tree](https://leetcode.com/problems/diameter-of-binary-tree/)** — `B`
- [ ] **268. [Merge Two Binary Trees](https://leetcode.com/problems/merge-two-binary-trees/)** — `B`
- [ ] **269. [Path Sum](https://leetcode.com/problems/path-sum/)** — `B`
- [ ] **270. [Count Complete Tree Nodes](https://leetcode.com/problems/count-complete-tree-nodes/)** — `B`
- [ ] **271. [Search in a Binary Search Tree](https://leetcode.com/problems/search-in-a-binary-search-tree/)** — `B`
- [ ] **272. [Insert into a Binary Search Tree](https://leetcode.com/problems/insert-into-a-binary-search-tree/)** — `B`
- [ ] **273. [Validate Binary Search Tree](https://leetcode.com/problems/validate-binary-search-tree/)** — `B`
- [ ] **274. [Convert Sorted Array to Binary Search Tree](https://leetcode.com/problems/convert-sorted-array-to-binary-search-tree/)** — `B`
- [ ] **275. [Kth Smallest Element in a BST](https://leetcode.com/problems/kth-smallest-element-in-a-bst/)** — `T`
- [ ] **276. [Lowest Common Ancestor of a Binary Search Tree](https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-search-tree/)** — `T`
- [ ] **277. [Binary Tree Level Order Traversal](https://leetcode.com/problems/binary-tree-level-order-traversal/)** — `T`
- [ ] **278. [Binary Tree Right Side View](https://leetcode.com/problems/binary-tree-right-side-view/)** — `T`
- [ ] **279. [Average of Levels in Binary Tree](https://leetcode.com/problems/average-of-levels-in-binary-tree/)** — `T`
- [ ] **280. [Binary Tree Zigzag Level Order Traversal](https://leetcode.com/problems/binary-tree-zigzag-level-order-traversal/)** — `T`

**Exit check:** pick 3 problems from Foundation, 2 from Build, and 1 from Transfer at random. Before coding each, state the invariant/pattern in one sentence.

---

## 15 — Advanced Trees, Paths & Reconstruction

**Goal:** Handle tree reconstruction, path aggregation, LCA, serialization, and harder subtree contracts.

**Recognition cue:** *What information must each subtree return upward for the parent to finish its job?*

- [ ] **281. [Binary Tree Paths](https://leetcode.com/problems/binary-tree-paths/)** — `F`
- [ ] **282. [Sum Root to Leaf Numbers](https://leetcode.com/problems/sum-root-to-leaf-numbers/)** — `F`
- [ ] **283. [Path Sum II](https://leetcode.com/problems/path-sum-ii/)** — `F`
- [ ] **284. [Path Sum III](https://leetcode.com/problems/path-sum-iii/)** — `F`
- [ ] **285. [Lowest Common Ancestor of a Binary Tree](https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-tree/)** — `F`
- [ ] **286. [Construct Binary Tree from Preorder and Inorder Traversal](https://leetcode.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal/)** — `F`
- [ ] **287. [Construct Binary Tree from Inorder and Postorder Traversal](https://leetcode.com/problems/construct-binary-tree-from-inorder-and-postorder-traversal/)** — `B`
- [ ] **288. [Flatten Binary Tree to Linked List](https://leetcode.com/problems/flatten-binary-tree-to-linked-list/)** — `B`
- [ ] **289. [Populating Next Right Pointers in Each Node](https://leetcode.com/problems/populating-next-right-pointers-in-each-node/)** — `B`
- [ ] **290. [Binary Search Tree Iterator](https://leetcode.com/problems/binary-search-tree-iterator/)** — `B`
- [ ] **291. [Delete Node in a BST](https://leetcode.com/problems/delete-node-in-a-bst/)** — `B`
- [ ] **292. [Recover Binary Search Tree](https://leetcode.com/problems/recover-binary-search-tree/)** — `B`
- [ ] **293. [House Robber III](https://leetcode.com/problems/house-robber-iii/)** — `B`
- [ ] **294. [All Nodes Distance K in Binary Tree](https://leetcode.com/problems/all-nodes-distance-k-in-binary-tree/)** — `B`
- [ ] **295. [Smallest Subtree with all the Deepest Nodes](https://leetcode.com/problems/smallest-subtree-with-all-the-deepest-nodes/)** — `T`
- [ ] **296. [Vertical Order Traversal of a Binary Tree](https://leetcode.com/problems/vertical-order-traversal-of-a-binary-tree/)** — `T`
- [ ] **297. [Boundary of Binary Tree](https://leetcode.com/problems/boundary-of-binary-tree/)** — `T`
- [ ] **298. [Serialize and Deserialize Binary Tree](https://leetcode.com/problems/serialize-and-deserialize-binary-tree/)** — `T`
- [ ] **299. [Binary Tree Maximum Path Sum](https://leetcode.com/problems/binary-tree-maximum-path-sum/)** — `T`
- [ ] **300. [Binary Tree Cameras](https://leetcode.com/problems/binary-tree-cameras/)** — `T`

**Exit check:** pick 3 problems from Foundation, 2 from Build, and 1 from Transfer at random. Before coding each, state the invariant/pattern in one sentence.

---

## 16 — Heaps & Priority Queues

**Goal:** Use heaps when you repeatedly need the smallest/largest K or the next best item.

**Recognition cue:** *Do I need repeated access to an extreme while data keeps changing?*

- [ ] **301. [Last Stone Weight](https://leetcode.com/problems/last-stone-weight/)** — `F`
- [ ] **302. [Kth Largest Element in a Stream](https://leetcode.com/problems/kth-largest-element-in-a-stream/)** — `F`
- [ ] **303. [Kth Largest Element in an Array](https://leetcode.com/problems/kth-largest-element-in-an-array/)** — `F`
- [ ] **304. [K Closest Points to Origin](https://leetcode.com/problems/k-closest-points-to-origin/)** — `F`
- [ ] **305. [Top K Frequent Words](https://leetcode.com/problems/top-k-frequent-words/)** — `F`
- [ ] **306. [Seat Reservation Manager](https://leetcode.com/problems/seat-reservation-manager/)** — `F`
- [ ] **307. [Smallest Number in Infinite Set](https://leetcode.com/problems/smallest-number-in-infinite-set/)** — `B`
- [ ] **308. [Task Scheduler](https://leetcode.com/problems/task-scheduler/)** — `B`
- [ ] **309. [Reorganize String](https://leetcode.com/problems/reorganize-string/)** — `B`
- [ ] **310. [Merge k Sorted Lists](https://leetcode.com/problems/merge-k-sorted-lists/)** — `B`
- [ ] **311. [Find Median from Data Stream](https://leetcode.com/problems/find-median-from-data-stream/)** — `B`
- [ ] **312. [Ugly Number II](https://leetcode.com/problems/ugly-number-ii/)** — `B`
- [ ] **313. [Furthest Building You Can Reach](https://leetcode.com/problems/furthest-building-you-can-reach/)** — `B`
- [ ] **314. [Single-Threaded CPU](https://leetcode.com/problems/single-threaded-cpu/)** — `B`
- [ ] **315. [Meeting Rooms III](https://leetcode.com/problems/meeting-rooms-iii/)** — `T`
- [ ] **316. [Total Cost to Hire K Workers](https://leetcode.com/problems/total-cost-to-hire-k-workers/)** — `T`
- [ ] **317. [Maximum Subsequence Score](https://leetcode.com/problems/maximum-subsequence-score/)** — `T`
- [ ] **318. [IPO](https://leetcode.com/problems/ipo/)** — `T`
- [ ] **319. [Smallest Range Covering Elements from K Lists](https://leetcode.com/problems/smallest-range-covering-elements-from-k-lists/)** — `T`
- [ ] **320. [Design Twitter](https://leetcode.com/problems/design-twitter/)** — `T`

**Exit check:** pick 3 problems from Foundation, 2 from Build, and 1 from Transfer at random. Before coding each, state the invariant/pattern in one sentence.

---

## 17 — Greedy

**Goal:** Learn when a locally optimal choice can be proven safe.

**Recognition cue:** *What choice can I commit to now without needing to reconsider it later?*

- [ ] **321. [Assign Cookies](https://leetcode.com/problems/assign-cookies/)** — `F`
- [ ] **322. [Lemonade Change](https://leetcode.com/problems/lemonade-change/)** — `F`
- [ ] **323. [Can Place Flowers](https://leetcode.com/problems/can-place-flowers/)** — `F`
- [ ] **324. [Maximum Units on a Truck](https://leetcode.com/problems/maximum-units-on-a-truck/)** — `F`
- [ ] **325. [Minimum Cost to Move Chips to The Same Position](https://leetcode.com/problems/minimum-cost-to-move-chips-to-the-same-position/)** — `F`
- [ ] **326. [Best Time to Buy and Sell Stock II](https://leetcode.com/problems/best-time-to-buy-and-sell-stock-ii/)** — `F`
- [ ] **327. [Jump Game](https://leetcode.com/problems/jump-game/)** — `B`
- [ ] **328. [Jump Game II](https://leetcode.com/problems/jump-game-ii/)** — `B`
- [ ] **329. [Gas Station](https://leetcode.com/problems/gas-station/)** — `B`
- [ ] **330. [Partition Labels](https://leetcode.com/problems/partition-labels/)** — `B`
- [ ] **331. [Hand of Straights](https://leetcode.com/problems/hand-of-straights/)** — `B`
- [ ] **332. [Merge Triplets to Form Target Triplet](https://leetcode.com/problems/merge-triplets-to-form-target-triplet/)** — `B`
- [ ] **333. [Valid Parenthesis String](https://leetcode.com/problems/valid-parenthesis-string/)** — `B`
- [ ] **334. [Bag of Tokens](https://leetcode.com/problems/bag-of-tokens/)** — `B`
- [ ] **335. [Eliminate Maximum Number of Monsters](https://leetcode.com/problems/eliminate-maximum-number-of-monsters/)** — `T`
- [ ] **336. [Two City Scheduling](https://leetcode.com/problems/two-city-scheduling/)** — `T`
- [ ] **337. [Broken Calculator](https://leetcode.com/problems/broken-calculator/)** — `T`
- [ ] **338. [Maximum Swap](https://leetcode.com/problems/maximum-swap/)** — `T`
- [ ] **339. [Monotone Increasing Digits](https://leetcode.com/problems/monotone-increasing-digits/)** — `T`
- [ ] **340. [Earliest Possible Day of Full Bloom](https://leetcode.com/problems/earliest-possible-day-of-full-bloom/)** — `T`

**Exit check:** pick 3 problems from Foundation, 2 from Build, and 1 from Transfer at random. Before coding each, state the invariant/pattern in one sentence.

---

## 18 — Graphs: BFS & DFS

**Goal:** Make graph traversal, visited-state management, components, and multi-source BFS routine.

**Recognition cue:** *What are the nodes, what are the edges, and when is a state considered visited?*

- [ ] **341. [Flood Fill](https://leetcode.com/problems/flood-fill/)** — `F`
- [ ] **342. [Island Perimeter](https://leetcode.com/problems/island-perimeter/)** — `F`
- [ ] **343. [Find if Path Exists in Graph](https://leetcode.com/problems/find-if-path-exists-in-graph/)** — `F`
- [ ] **344. [Keys and Rooms](https://leetcode.com/problems/keys-and-rooms/)** — `F`
- [ ] **345. [Number of Provinces](https://leetcode.com/problems/number-of-provinces/)** — `F`
- [ ] **346. [Number of Islands](https://leetcode.com/problems/number-of-islands/)** — `F`
- [ ] **347. [Max Area of Island](https://leetcode.com/problems/max-area-of-island/)** — `B`
- [ ] **348. [Surrounded Regions](https://leetcode.com/problems/surrounded-regions/)** — `B`
- [ ] **349. [Clone Graph](https://leetcode.com/problems/clone-graph/)** — `B`
- [ ] **350. [All Paths From Source to Target](https://leetcode.com/problems/all-paths-from-source-to-target/)** — `B`
- [ ] **351. [Pacific Atlantic Water Flow](https://leetcode.com/problems/pacific-atlantic-water-flow/)** — `B`
- [ ] **352. [Number of Enclaves](https://leetcode.com/problems/number-of-enclaves/)** — `B`
- [ ] **353. [Count Sub Islands](https://leetcode.com/problems/count-sub-islands/)** — `B`
- [ ] **354. [01 Matrix](https://leetcode.com/problems/01-matrix/)** — `B`
- [ ] **355. [As Far from Land as Possible](https://leetcode.com/problems/as-far-from-land-as-possible/)** — `T`
- [ ] **356. [Nearest Exit from Entrance in Maze](https://leetcode.com/problems/nearest-exit-from-entrance-in-maze/)** — `T`
- [ ] **357. [Shortest Path in Binary Matrix](https://leetcode.com/problems/shortest-path-in-binary-matrix/)** — `T`
- [ ] **358. [Minimum Genetic Mutation](https://leetcode.com/problems/minimum-genetic-mutation/)** — `T`
- [ ] **359. [Word Ladder](https://leetcode.com/problems/word-ladder/)** — `T`
- [ ] **360. [Shortest Bridge](https://leetcode.com/problems/shortest-bridge/)** — `T`

**Exit check:** pick 3 problems from Foundation, 2 from Build, and 1 from Transfer at random. Before coding each, state the invariant/pattern in one sentence.

---

## 19 — Union-Find, DAGs & Topological Sort

**Goal:** Learn connectivity merging, cycle detection, dependency ordering, and DAG reasoning.

**Recognition cue:** *Is this about components (DSU) or prerequisites/indegree (topological sort)?*

- [ ] **361. [Redundant Connection](https://leetcode.com/problems/redundant-connection/)** — `F`
- [ ] **362. [Graph Valid Tree](https://leetcode.com/problems/graph-valid-tree/)** — `F`
- [ ] **363. [Number of Connected Components in an Undirected Graph](https://leetcode.com/problems/number-of-connected-components-in-an-undirected-graph/)** — `F`
- [ ] **364. [Accounts Merge](https://leetcode.com/problems/accounts-merge/)** — `F`
- [ ] **365. [Satisfiability of Equality Equations](https://leetcode.com/problems/satisfiability-of-equality-equations/)** — `F`
- [ ] **366. [Most Stones Removed with Same Row or Column](https://leetcode.com/problems/most-stones-removed-with-same-row-or-column/)** — `F`
- [ ] **367. [Regions Cut By Slashes](https://leetcode.com/problems/regions-cut-by-slashes/)** — `B`
- [ ] **368. [Smallest String With Swaps](https://leetcode.com/problems/smallest-string-with-swaps/)** — `B`
- [ ] **369. [Similar String Groups](https://leetcode.com/problems/similar-string-groups/)** — `B`
- [ ] **370. [Course Schedule](https://leetcode.com/problems/course-schedule/)** — `B`
- [ ] **371. [Course Schedule II](https://leetcode.com/problems/course-schedule-ii/)** — `B`
- [ ] **372. [Find Eventual Safe States](https://leetcode.com/problems/find-eventual-safe-states/)** — `B`
- [ ] **373. [Minimum Height Trees](https://leetcode.com/problems/minimum-height-trees/)** — `B`
- [ ] **374. [Alien Dictionary](https://leetcode.com/problems/alien-dictionary/)** — `B`
- [ ] **375. [Parallel Courses](https://leetcode.com/problems/parallel-courses/)** — `T`
- [ ] **376. [Sequence Reconstruction](https://leetcode.com/problems/sequence-reconstruction/)** — `T`
- [ ] **377. [Find All Possible Recipes from Given Supplies](https://leetcode.com/problems/find-all-possible-recipes-from-given-supplies/)** — `T`
- [ ] **378. [Sort Items by Groups Respecting Dependencies](https://leetcode.com/problems/sort-items-by-groups-respecting-dependencies/)** — `T`
- [ ] **379. [Largest Color Value in a Directed Graph](https://leetcode.com/problems/largest-color-value-in-a-directed-graph/)** — `T`
- [ ] **380. [Min Cost to Connect All Points](https://leetcode.com/problems/min-cost-to-connect-all-points/)** — `T`

**Exit check:** pick 3 problems from Foundation, 2 from Build, and 1 from Transfer at random. Before coding each, state the invariant/pattern in one sentence.

---

## 20 — Shortest Paths, Weighted Graphs & MST

**Goal:** Build weighted-graph instincts: Dijkstra, Bellman-Ford variants, state-expanded shortest paths, and MST.

**Recognition cue:** *What is the state, what is the edge cost, and is greedy shortest-path finalization valid?*

- [ ] **381. [Network Delay Time](https://leetcode.com/problems/network-delay-time/)** — `F`
- [ ] **382. [Cheapest Flights Within K Stops](https://leetcode.com/problems/cheapest-flights-within-k-stops/)** — `F`
- [ ] **383. [Path With Minimum Effort](https://leetcode.com/problems/path-with-minimum-effort/)** — `F`
- [ ] **384. [Swim in Rising Water](https://leetcode.com/problems/swim-in-rising-water/)** — `F`
- [ ] **385. [Path with Maximum Probability](https://leetcode.com/problems/path-with-maximum-probability/)** — `F`
- [ ] **386. [Minimum Cost to Make at Least One Valid Path in a Grid](https://leetcode.com/problems/minimum-cost-to-make-at-least-one-valid-path-in-a-grid/)** — `F`
- [ ] **387. [Minimum Obstacle Removal to Reach Corner](https://leetcode.com/problems/minimum-obstacle-removal-to-reach-corner/)** — `B`
- [ ] **388. [Reachable Nodes In Subdivided Graph](https://leetcode.com/problems/reachable-nodes-in-subdivided-graph/)** — `B`
- [ ] **389. [Number of Ways to Arrive at Destination](https://leetcode.com/problems/number-of-ways-to-arrive-at-destination/)** — `B`
- [ ] **390. [Find the City With the Smallest Number of Neighbors at a Threshold Distance](https://leetcode.com/problems/find-the-city-with-the-smallest-number-of-neighbors-at-a-threshold-distance/)** — `B`
- [ ] **391. [Evaluate Division](https://leetcode.com/problems/evaluate-division/)** — `B`
- [ ] **392. [The Maze II](https://leetcode.com/problems/the-maze-ii/)** — `B`
- [ ] **393. [The Maze III](https://leetcode.com/problems/the-maze-iii/)** — `B`
- [ ] **394. [Path With Maximum Minimum Value](https://leetcode.com/problems/path-with-maximum-minimum-value/)** — `B`
- [ ] **395. [Minimum Cost to Reach City With Discounts](https://leetcode.com/problems/minimum-cost-to-reach-city-with-discounts/)** — `T`
- [ ] **396. [Connecting Cities With Minimum Cost](https://leetcode.com/problems/connecting-cities-with-minimum-cost/)** — `T`
- [ ] **397. [Optimize Water Distribution in a Village](https://leetcode.com/problems/optimize-water-distribution-in-a-village/)** — `T`
- [ ] **398. [Critical Connections in a Network](https://leetcode.com/problems/critical-connections-in-a-network/)** — `T`
- [ ] **399. [Reconstruct Itinerary](https://leetcode.com/problems/reconstruct-itinerary/)** — `T`
- [ ] **400. [Shortest Path Visiting All Nodes](https://leetcode.com/problems/shortest-path-visiting-all-nodes/)** — `T`

**Exit check:** pick 3 problems from Foundation, 2 from Build, and 1 from Transfer at random. Before coding each, state the invariant/pattern in one sentence.

---

## 21 — Dynamic Programming I: 1D State

**Goal:** Learn to define compact 1D DP state and transitions before touching code.

**Recognition cue:** *What does dp[i] mean in one sentence, and which previous states can produce it?*

- [ ] **401. [Climbing Stairs](https://leetcode.com/problems/climbing-stairs/)** — `F`
- [ ] **402. [Min Cost Climbing Stairs](https://leetcode.com/problems/min-cost-climbing-stairs/)** — `F`
- [ ] **403. [N-th Tribonacci Number](https://leetcode.com/problems/n-th-tribonacci-number/)** — `F`
- [ ] **404. [House Robber](https://leetcode.com/problems/house-robber/)** — `F`
- [ ] **405. [House Robber II](https://leetcode.com/problems/house-robber-ii/)** — `F`
- [ ] **406. [Delete and Earn](https://leetcode.com/problems/delete-and-earn/)** — `F`
- [ ] **407. [Perfect Squares](https://leetcode.com/problems/perfect-squares/)** — `B`
- [ ] **408. [Decode Ways](https://leetcode.com/problems/decode-ways/)** — `B`
- [ ] **409. [Word Break](https://leetcode.com/problems/word-break/)** — `B`
- [ ] **410. [Integer Break](https://leetcode.com/problems/integer-break/)** — `B`
- [ ] **411. [Coin Change](https://leetcode.com/problems/coin-change/)** — `B`
- [ ] **412. [Combination Sum IV](https://leetcode.com/problems/combination-sum-iv/)** — `B`
- [ ] **413. [Maximum Product Subarray](https://leetcode.com/problems/maximum-product-subarray/)** — `B`
- [ ] **414. [Longest Increasing Subsequence](https://leetcode.com/problems/longest-increasing-subsequence/)** — `B`
- [ ] **415. [Number of Longest Increasing Subsequence](https://leetcode.com/problems/number-of-longest-increasing-subsequence/)** — `T`
- [ ] **416. [Wiggle Subsequence](https://leetcode.com/problems/wiggle-subsequence/)** — `T`
- [ ] **417. [Best Time to Buy and Sell Stock with Cooldown](https://leetcode.com/problems/best-time-to-buy-and-sell-stock-with-cooldown/)** — `T`
- [ ] **418. [Best Time to Buy and Sell Stock with Transaction Fee](https://leetcode.com/problems/best-time-to-buy-and-sell-stock-with-transaction-fee/)** — `T`
- [ ] **419. [Maximum Sum Circular Subarray](https://leetcode.com/problems/maximum-sum-circular-subarray/)** — `T`
- [ ] **420. [Partition Equal Subset Sum](https://leetcode.com/problems/partition-equal-subset-sum/)** — `T`

**Exit check:** pick 3 problems from Foundation, 2 from Build, and 1 from Transfer at random. Before coding each, state the invariant/pattern in one sentence.

---

## 22 — Dynamic Programming II: Grids & Knapsack

**Goal:** Extend DP to grids, capacities, and two-dimensional state spaces.

**Recognition cue:** *Which two independent variables define the subproblem, and in what order can states be filled?*

- [ ] **421. [Unique Paths](https://leetcode.com/problems/unique-paths/)** — `F`
- [ ] **422. [Unique Paths II](https://leetcode.com/problems/unique-paths-ii/)** — `F`
- [ ] **423. [Minimum Path Sum](https://leetcode.com/problems/minimum-path-sum/)** — `F`
- [ ] **424. [Triangle](https://leetcode.com/problems/triangle/)** — `F`
- [ ] **425. [Minimum Falling Path Sum](https://leetcode.com/problems/minimum-falling-path-sum/)** — `F`
- [ ] **426. [Minimum Falling Path Sum II](https://leetcode.com/problems/minimum-falling-path-sum-ii/)** — `F`
- [ ] **427. [Dungeon Game](https://leetcode.com/problems/dungeon-game/)** — `B`
- [ ] **428. [Maximal Square](https://leetcode.com/problems/maximal-square/)** — `B`
- [ ] **429. [Longest Increasing Path in a Matrix](https://leetcode.com/problems/longest-increasing-path-in-a-matrix/)** — `B`
- [ ] **430. [Out of Boundary Paths](https://leetcode.com/problems/out-of-boundary-paths/)** — `B`
- [ ] **431. [Knight Probability in Chessboard](https://leetcode.com/problems/knight-probability-in-chessboard/)** — `B`
- [ ] **432. [Coin Change II](https://leetcode.com/problems/coin-change-ii/)** — `B`
- [ ] **433. [Target Sum](https://leetcode.com/problems/target-sum/)** — `B`
- [ ] **434. [Ones and Zeroes](https://leetcode.com/problems/ones-and-zeroes/)** — `B`
- [ ] **435. [Last Stone Weight II](https://leetcode.com/problems/last-stone-weight-ii/)** — `T`
- [ ] **436. [Profitable Schemes](https://leetcode.com/problems/profitable-schemes/)** — `T`
- [ ] **437. [Paint House](https://leetcode.com/problems/paint-house/)** — `T`
- [ ] **438. [Paint House II](https://leetcode.com/problems/paint-house-ii/)** — `T`
- [ ] **439. [Cherry Pickup II](https://leetcode.com/problems/cherry-pickup-ii/)** — `T`
- [ ] **440. [Cherry Pickup](https://leetcode.com/problems/cherry-pickup/)** — `T`

**Exit check:** pick 3 problems from Foundation, 2 from Build, and 1 from Transfer at random. Before coding each, state the invariant/pattern in one sentence.

---

## 23 — Dynamic Programming III: Strings, Sequences & Games

**Goal:** Handle sequence alignment, palindrome intervals, string DP, and game/interval DP.

**Recognition cue:** *Is the state a prefix pair, an interval [l,r], or a turn/choice over a smaller interval?*

- [ ] **441. [Longest Common Subsequence](https://leetcode.com/problems/longest-common-subsequence/)** — `F`
- [ ] **442. [Delete Operation for Two Strings](https://leetcode.com/problems/delete-operation-for-two-strings/)** — `F`
- [ ] **443. [Uncrossed Lines](https://leetcode.com/problems/uncrossed-lines/)** — `F`
- [ ] **444. [Minimum ASCII Delete Sum for Two Strings](https://leetcode.com/problems/minimum-ascii-delete-sum-for-two-strings/)** — `F`
- [ ] **445. [Edit Distance](https://leetcode.com/problems/edit-distance/)** — `F`
- [ ] **446. [Distinct Subsequences](https://leetcode.com/problems/distinct-subsequences/)** — `F`
- [ ] **447. [Interleaving String](https://leetcode.com/problems/interleaving-string/)** — `B`
- [ ] **448. [Palindromic Substrings](https://leetcode.com/problems/palindromic-substrings/)** — `B`
- [ ] **449. [Longest Palindromic Subsequence](https://leetcode.com/problems/longest-palindromic-subsequence/)** — `B`
- [ ] **450. [Longest Palindromic Substring](https://leetcode.com/problems/longest-palindromic-substring/)** — `B`
- [ ] **451. [Minimum Insertion Steps to Make a String Palindrome](https://leetcode.com/problems/minimum-insertion-steps-to-make-a-string-palindrome/)** — `B`
- [ ] **452. [Shortest Common Supersequence](https://leetcode.com/problems/shortest-common-supersequence/)** — `B`
- [ ] **453. [Palindrome Partitioning II](https://leetcode.com/problems/palindrome-partitioning-ii/)** — `B`
- [ ] **454. [Regular Expression Matching](https://leetcode.com/problems/regular-expression-matching/)** — `B`
- [ ] **455. [Wildcard Matching](https://leetcode.com/problems/wildcard-matching/)** — `T`
- [ ] **456. [Scramble String](https://leetcode.com/problems/scramble-string/)** — `T`
- [ ] **457. [Predict the Winner](https://leetcode.com/problems/predict-the-winner/)** — `T`
- [ ] **458. [Stone Game](https://leetcode.com/problems/stone-game/)** — `T`
- [ ] **459. [Stone Game II](https://leetcode.com/problems/stone-game-ii/)** — `T`
- [ ] **460. [Burst Balloons](https://leetcode.com/problems/burst-balloons/)** — `T`

**Exit check:** pick 3 problems from Foundation, 2 from Build, and 1 from Transfer at random. Before coding each, state the invariant/pattern in one sentence.

---

## 24 — Tries & Bit Manipulation

**Goal:** Learn prefix trees and bit-level identities used in interviews.

**Recognition cue:** *Is the structure driven by prefixes, or can XOR/bit masks encode the state more directly?*

- [ ] **461. [Implement Trie (Prefix Tree)](https://leetcode.com/problems/implement-trie-prefix-tree/)** — `F`
- [ ] **462. [Design Add and Search Words Data Structure](https://leetcode.com/problems/design-add-and-search-words-data-structure/)** — `F`
- [ ] **463. [Replace Words](https://leetcode.com/problems/replace-words/)** — `F`
- [ ] **464. [Map Sum Pairs](https://leetcode.com/problems/map-sum-pairs/)** — `F`
- [ ] **465. [Longest Word in Dictionary](https://leetcode.com/problems/longest-word-in-dictionary/)** — `F`
- [ ] **466. [Search Suggestions System](https://leetcode.com/problems/search-suggestions-system/)** — `F`
- [ ] **467. [Word Search II](https://leetcode.com/problems/word-search-ii/)** — `B`
- [ ] **468. [Concatenated Words](https://leetcode.com/problems/concatenated-words/)** — `B`
- [ ] **469. [Stream of Characters](https://leetcode.com/problems/stream-of-characters/)** — `B`
- [ ] **470. [Single Number](https://leetcode.com/problems/single-number/)** — `B`
- [ ] **471. [Number of 1 Bits](https://leetcode.com/problems/number-of-1-bits/)** — `B`
- [ ] **472. [Counting Bits](https://leetcode.com/problems/counting-bits/)** — `B`
- [ ] **473. [Reverse Bits](https://leetcode.com/problems/reverse-bits/)** — `B`
- [ ] **474. [Missing Number](https://leetcode.com/problems/missing-number/)** — `B`
- [ ] **475. [Power of Two](https://leetcode.com/problems/power-of-two/)** — `T`
- [ ] **476. [Single Number II](https://leetcode.com/problems/single-number-ii/)** — `T`
- [ ] **477. [Single Number III](https://leetcode.com/problems/single-number-iii/)** — `T`
- [ ] **478. [Sum of Two Integers](https://leetcode.com/problems/sum-of-two-integers/)** — `T`
- [ ] **479. [Bitwise AND of Numbers Range](https://leetcode.com/problems/bitwise-and-of-numbers-range/)** — `T`
- [ ] **480. [Maximum XOR of Two Numbers in an Array](https://leetcode.com/problems/maximum-xor-of-two-numbers-in-an-array/)** — `T`

**Exit check:** pick 3 problems from Foundation, 2 from Build, and 1 from Transfer at random. Before coding each, state the invariant/pattern in one sentence.

---

## 25 — Math, Geometry & Matrix Manipulation

**Goal:** Cover arithmetic, coordinate geometry, matrix transforms, and randomized selection.

**Recognition cue:** *What mathematical invariant or coordinate transform removes the need for simulation?*

- [ ] **481. [Palindrome Number](https://leetcode.com/problems/palindrome-number/)** — `F`
- [ ] **482. [Roman to Integer](https://leetcode.com/problems/roman-to-integer/)** — `F`
- [ ] **483. [Integer to Roman](https://leetcode.com/problems/integer-to-roman/)** — `F`
- [ ] **484. [Excel Sheet Column Number](https://leetcode.com/problems/excel-sheet-column-number/)** — `F`
- [ ] **485. [Excel Sheet Column Title](https://leetcode.com/problems/excel-sheet-column-title/)** — `F`
- [ ] **486. [Factorial Trailing Zeroes](https://leetcode.com/problems/factorial-trailing-zeroes/)** — `F`
- [ ] **487. [Add Digits](https://leetcode.com/problems/add-digits/)** — `B`
- [ ] **488. [Power of Three](https://leetcode.com/problems/power-of-three/)** — `B`
- [ ] **489. [Ugly Number](https://leetcode.com/problems/ugly-number/)** — `B`
- [ ] **490. [Greatest Common Divisor of Strings](https://leetcode.com/problems/greatest-common-divisor-of-strings/)** — `B`
- [ ] **491. [Count Primes](https://leetcode.com/problems/count-primes/)** — `B`
- [ ] **492. [Spiral Matrix II](https://leetcode.com/problems/spiral-matrix-ii/)** — `B`
- [ ] **493. [Rotate Image](https://leetcode.com/problems/rotate-image/)** — `B`
- [ ] **494. [Valid Sudoku](https://leetcode.com/problems/valid-sudoku/)** — `B`
- [ ] **495. [Rectangle Overlap](https://leetcode.com/problems/rectangle-overlap/)** — `T`
- [ ] **496. [Rectangle Area](https://leetcode.com/problems/rectangle-area/)** — `T`
- [ ] **497. [Max Points on a Line](https://leetcode.com/problems/max-points-on-a-line/)** — `T`
- [ ] **498. [Detect Squares](https://leetcode.com/problems/detect-squares/)** — `T`
- [ ] **499. [Random Pick with Weight](https://leetcode.com/problems/random-pick-with-weight/)** — `T`
- [ ] **500. [Random Point in Non-overlapping Rectangles](https://leetcode.com/problems/random-point-in-non-overlapping-rectangles/)** — `T`

**Exit check:** pick 3 problems from Foundation, 2 from Build, and 1 from Transfer at random. Before coding each, state the invariant/pattern in one sentence.

---

## 26 — Design Problems & OA Hardening

**Goal:** Practice composing multiple data structures behind a clean API—the closest LC gets to small production components.

**Recognition cue:** *What operations must be O(1)/O(log n), and which combination of structures gives those guarantees?*

- [ ] **501. [Design HashSet](https://leetcode.com/problems/design-hashset/)** — `F`
- [ ] **502. [Design HashMap](https://leetcode.com/problems/design-hashmap/)** — `F`
- [ ] **503. [Design Linked List](https://leetcode.com/problems/design-linked-list/)** — `F`
- [ ] **504. [Insert Delete GetRandom O(1)](https://leetcode.com/problems/insert-delete-getrandom-o1/)** — `F`
- [ ] **505. [Insert Delete GetRandom O(1) - Duplicates allowed](https://leetcode.com/problems/insert-delete-getrandom-o1-duplicates-allowed/)** — `F`
- [ ] **506. [LRU Cache](https://leetcode.com/problems/lru-cache/)** — `F`
- [ ] **507. [LFU Cache](https://leetcode.com/problems/lfu-cache/)** — `B`
- [ ] **508. [Snapshot Array](https://leetcode.com/problems/snapshot-array/)** — `B`
- [ ] **509. [Design Browser History](https://leetcode.com/problems/design-browser-history/)** — `B`
- [ ] **510. [Design Underground System](https://leetcode.com/problems/design-underground-system/)** — `B`
- [ ] **511. [Logger Rate Limiter](https://leetcode.com/problems/logger-rate-limiter/)** — `B`
- [ ] **512. [Design In-Memory File System](https://leetcode.com/problems/design-in-memory-file-system/)** — `B`
- [ ] **513. [Design Tic-Tac-Toe](https://leetcode.com/problems/design-tic-tac-toe/)** — `B`
- [ ] **514. [Design Snake Game](https://leetcode.com/problems/design-snake-game/)** — `B`
- [ ] **515. [Design Excel Sum Formula](https://leetcode.com/problems/design-excel-sum-formula/)** — `T`
- [ ] **516. [Design Search Autocomplete System](https://leetcode.com/problems/design-search-autocomplete-system/)** — `T`
- [ ] **517. [Encode and Decode TinyURL](https://leetcode.com/problems/encode-and-decode-tinyurl/)** — `T`
- [ ] **518. [All O`one Data Structure](https://leetcode.com/problems/all-oone-data-structure/)** — `T`
- [ ] **519. [Design Authentication Manager](https://leetcode.com/problems/design-authentication-manager/)** — `T`
- [ ] **520. [Design a Food Rating System](https://leetcode.com/problems/design-a-food-rating-system/)** — `T`

**Exit check:** pick 3 problems from Foundation, 2 from Build, and 1 from Transfer at random. Before coding each, state the invariant/pattern in one sentence.

---

## Suggested review cadence

| When | What to do |
|---|---|
| Same day | Solve/learn the problem; rebuild from blank after reading help. |
| +1 day | Re-solve 2–4 misses from memory. |
| +3 days | One random problem from the current module. |
| +7 days | Timed retention pass. |
| End of module | Six-problem random gate. |
| Every 4 modules | 60–90 minute mixed OA set from completed modules. |

## What *not* to do

- Do not browse raw LeetCode tags looking for a learning sequence.
- Do not spend an hour proving to yourself that you cannot discover a known trick.
- Do not memorize code line-by-line; memorize **state + invariant + transition**.
- Do not optimize a correct brute-force idea before you can explain why the optimization works.
- Do not count a problem as learned just because you understood the editorial while reading it.

## OA endgame

After finishing Modules 1–20, start one mixed timed set per week while continuing the DP/trie/design modules. After all 26 modules, spend 4–6 weeks mostly on mixed sets and targeted weak-pattern revision rather than adding random new questions.

Your final objective is not to remember 520 solutions. It is to compress them into a few dozen reusable mental models.
